package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"

	log "github.com/Sirupsen/logrus"
	"strconv"
	"database/sql"
	"os"
	"encoding/json"
)

const (
	xMax           = 500
	yMax           = 500
	xMin           = 0
	yMin           = 0
	levelUpSeconds = 600
	levelUpBase    = float64(1.16)
)

type Game struct {
	startedAt        time.Time
	heroes           []HeroDB
	adminToken       string
	joinChan         chan JoinRequest
	activateHeroChan chan ActivateHeroRequest
	exitChan         chan []byte
	dbconnection	 string
}


//Db_Configuration struct holds the fields for the Database Connection string
type Db_configuration struct {

	Db_user	string 	`json:"Dbuser"`
	Db_pass	string 	`json:"Dbpass"`
	Db_ip	string 	`json:"Dbip"`
	Db_port	string	`json:"Dbport"`
	Db_name	string 	`json:"Dbname"`
}


// Db_connection returns the database connection string
// from a json configuration file
//Parameters: None
//Return: String with the DB Connection String
func Db_connection() string {
	// Returns the Connection String for the Database

	file,err := os.Open("dbconfig.json")

	if err != nil {
		fmt.Println("[Engine] : Db_connection : File Error:", err)
	}

	decoder := json.NewDecoder(file)

	var db_conf = Db_configuration{}

	err2 := decoder.Decode(&db_conf)

	if err2 != nil {
		fmt.Println("[Engine] : Db_connection : Decoder Error:", err2)
	}

	return db_conf.Db_user + ":" + db_conf.Db_pass + "@tcp(" + db_conf.Db_ip + ":" + db_conf.Db_port + ")/" + db_conf.Db_name

}
// NewGame creates a new game
func NewGame(adminToken string) *Game {

	//Get the Database Connection string
	var dbconn = Db_connection()

	game := &Game{
		startedAt:        time.Now(),
		heroes:           []HeroDB{},
		joinChan:         make(chan JoinRequest),
		activateHeroChan: make(chan ActivateHeroRequest),
		exitChan:         make(chan []byte),
		adminToken:       adminToken,
		dbconnection:	  dbconn,
	}
	return game
}

// StartGame starts the game
func StartGame(adminToken string) {
	// game := NewGame(adminToken)

	game, err := LoadFromDB()


	if err != nil {
		log.Panic(err)
	}

	go game.StartEngine()
	game.StartAPI()
}

// StartEngine starts the engine
func (g *Game) StartEngine() {

	g.dbconnection = Db_connection()
	log.Info("[Engine] StartEngine | DB Connection:" + g.dbconnection)

	ticker := time.NewTicker(time.Second * 2)
	tickerDB := time.NewTicker(time.Minute * 1)

	for {
		select {
		case <-ticker.C:
			g.moveHeroes()
			g.CheckHeroLevel()
		//TODO: check battles
		case <-tickerDB.C:
			log.Info("Saving game state to DB")
			if err := SaveToDB(g); err != nil {
				log.Error(err)
			}
		case req := <-g.joinChan:
			log.Info("Join hero")
			success, message := g.joinHero(req.name, req.email, req.heroClass, req.TokenRequest.token)
			req.Response <- GameResponse{success: success, message: message}
			close(req.Response)
		case req := <-g.activateHeroChan:
			log.Info("Activate hero")
			success := g.activateHero(req.name, req.TokenRequest.token)
			req.Response <- GameResponse{success: success, message: ""}
			close(req.Response)
		case <-g.exitChan:
			log.Info("Exiting game")
			return
		}
	}

}

func (g *Game) joinHero(name, email, class, adminToken string) (bool, string) {

	if !g.authorizeAdmin(adminToken) {
		return false, "You are not authorized to perform this action."
	}

	hero := &HeroDB{
		HeroName:        name,
		Email:       email,
		HClass:       class,
		Enabled:     false,
		Token:       randToken(),
		Level:       0,
		NextLevelAt: time.Now().Add(99999 * time.Hour),
		HeroCreatedAt:   time.Now(),
		Ring:     0,
		Amulet:   0,
		Charm:    0,
		Weapon:   0,
		Helm:     0,
		Tunic:    0,
		Gloves:   0,
		Shield:   0,
		Leggings: 0,
		Boots:    0,
		Xpos: rand.Intn(xMax-xMin) + xMin,
		Ypos: rand.Intn(yMax-yMin) + yMin,
	}

	g.heroes = append(g.heroes, *hero)
	log.Infof("Hero %s has been created, but will not play until it's activated.", hero.HeroName)
	return true, fmt.Sprintf("Token: %s", hero.Token)
}

func (g *Game) activateHero(name, token string) bool {

	i := g.getHeroIndex(name)
	if i == -1 {
		return false
	}
	if g.heroes[i].Token != token {
		return false
	}

	ttl := getTTL(1) // Time to level 1
	g.heroes[i].Enabled = true
	g.heroes[i].NextLevelAt = time.Now().Add(ttl * time.Second)
	log.Infof("Success! Hero %s has been activated and will reach level 1 in %d seconds.", g.heroes[i].HeroName, ttl)
	return true
}

func (g *Game) moveHeroes() {

	for i := range g.heroes {
		if !g.heroes[i].Enabled {
			continue
		}
		g.heroes[i].Xpos = truncateInt(g.heroes[i].Xpos+(rand.Intn(3)-1), xMin, xMax)
		g.heroes[i].Ypos = truncateInt(g.heroes[i].Ypos+(rand.Intn(3)-1), yMin, yMax)
	}
}

func (g *Game) authorizeAdmin(token string) bool {
	return g.adminToken == token
}

func (g *Game) getHeroIndex(name string) int {
	for i, hero := range g.heroes {
		if hero.HeroName == name {
			return i
		}
	}
	return -1
}

func getTTL(level int) time.Duration {
	return time.Duration(levelUpSeconds * (math.Pow(levelUpBase, float64(level))))
}


/*
//WORLD EVENTS

EVENT			Frequency
Hand of God 	20 hours
Team Battle		24 hours
Calamity 		8 hours
GodSend			4 hours
*/

//Hand of God function implements the Gods powers on the Realm. It happens 1 an hour and it has 1/4000 chances
// to strike a Hero. The outcome has 80 chances to be good and 20 chances to the bad
//Parameters: SQL DB connection Object
//Return: None
func (g *Game) HandOfGod(conn *sql.DB) {

	sqldb, err := sql.Open("mysql", g.dbconnection)

	if err != nil {
		log.Errorln("[Engine] : Hand of Good : Database settings failed: %s", err)
	}
	defer sqldb.Close()


	for i := range g.heroes {
		if !g.heroes[i].Enabled {
			continue
		}

		if(rand.Int31n(4000) == 1){

			var hero_chance = rand.Intn(10)
			var message string
			var hero_next_level = g.heroes[i].Level + 1
			var time_calculation = int((((rand.Intn(71)+5))/100) * hero_next_level * 3600)

			if(hero_chance >= 3 ) {
				// Good outcome

				message = "Verily I say undo thee, the Heavens have burst forth, and the blessed hand of God Carried" +
					g.heroes[i].HeroName + " for " + strconv.Itoa(time_calculation) + " seconds. Toward level " +
					strconv.Itoa(hero_next_level)

				//TODO: Update the Hero TTL for the next level

			} else { //Bad outcome


				message = "Thereupon He stretched out His little finger among them and consummed" +
					g.heroes[i].HeroName + " with fier, slowing the heathen " + strconv.Itoa(time_calculation) +
					" seconds from  level " + strconv.Itoa(hero_next_level)

				//TODO: Update the Hero TTL for the next level
			}

			log.Info("[Hand of God]" + message)

			//Add Event to WorldEvents and HeroWorldEvents Tables
			Insert_World_Event_for_Hero(g.heroes[i].HeroID, message, conn)


		}

	}
}

//God Send function implements the Gods gits to Heros  on the Realm. It happens 1 an hour and it has 1/4000 chances
// to strike a Hero. The outcome has 80 chances to be good and 20 chances to the bad
//Parameters: SQL DB connection Object
//Return: None
func (g *Game) GodSend() {

	//TODO: Complete GodSend
	/*
	  var good_events = map[int]string{}

		good_events[1] = " found a pair of nice Shoes"
		good_events[2] = " caught a Unicorn"
		good_events[3] = " discovered a secret, underground passage leading to Barcelona's best tavern"
		good_events[4] = " was taught to run quickly by a secret tribe of pygmies that know how to, among other things, run quickly"
		good_events[5] = " discovered caffeinated coffee"
		good_events[6] = " grew an extra leg"
		good_events[7] = " was visited by a very pretty nymph"
		good_events[8] = " found kitten"
		good_events[9] = " learned Python"
		good_events[10] = " found an exploit in the Titan Idle RPG code"
		good_events[11] = " tamed a wild horse"
		good_events[12] = " found a one-time-use spell of quickness"
		good_events[13] = " bought a faster computer"
		good_events[14] = " bribed the local OpenStack administrator"
		good_events[15] = " stopped using dial-up"
		good_events[16] = " invented the wheel"
		good_events[17] = " gained a sixth sense"
		good_events[18] = " got a kiss from drwiii"
		good_events[19] = " had his clothes laundered by a passing fairy"
		good_events[20] = " was rejuvenated by drinking from a magic stream"
		good_events[21] = " was bitten by a radioactive spider"
		good_events[22] = " hit it off with a drunk sorority chick named Jenny"
		good_events[23] = " was accepted into Pi Beta Phi"
		//OpenSTack Related
		good_events[24] = " was notified that Jenkins tests passed"
		good_events[25] = " got his first patch +4 approved in OpenStack"
		good_events[26] = " got a HEAT template successfully deployed"

	  items :=[10]string{"weapon","tunic","shield","leggins","ring","gloves","boots","helm","charm","amulet"}

	*/

	for i := range g.heroes {
		if !g.heroes[i].Enabled {
			continue
		}

		if(rand.Int31n(2000) == 1){

			if(rand.Intn(10) == 1) {  //Ultra Godsend


			} else {


			}
			g.heroes[i].Xpos = truncateInt(g.heroes[i].Xpos+(rand.Intn(3)-1), xMin, xMax)
			g.heroes[i].Ypos = truncateInt(g.heroes[i].Ypos+(rand.Intn(3)-1), yMin, yMax)

		}

	}
}

func (g *Game) Calamity() {


	sqldb, err := sql.Open("mysql", g.dbconnection)

	if err != nil {
		log.Errorln("[Engine] : Calamity : Database settings failed: %s", err)
	}
	defer sqldb.Close()


	//TODO Complete Calamity

	/*
		var bad_events = map[int]string{}

		bad_events[1] = " was bitten by Neutron"
		bad_events[2] = " fell into a hole"
		bad_events[3] = " bit their tongue"
		bad_events[4] = " set thyself on fire"
		bad_events[5] = " ate a poisonous fruit"
		bad_events[6] = " lost their mind"
		bad_events[7] = " died, temporarily.."
		bad_events[8] = " was caught in a terrible snowstorm"
		bad_events[9] = " EXPLODED, somewhat.."
		bad_events[10] = " got knifed in a dark alley"
		bad_events[11] = " saw an episode of Ally McBeal"
		bad_events[12] = " got turned INSIDE OUT, practically"
		bad_events[13] = " ate a very disagreeable fruit, getting a terrible case of heartburn"
		bad_events[14] = " met up with a mob hitman for not paying his hosting bills"
		bad_events[15] = " has fallen ill with the black plague"
		bad_events[16] = " was struck by lightning"
		bad_events[17] = " was attacked by a rabid giant rabbit"
		bad_events[18] = " was attacked by a rabid wolverine"
		bad_events[19] = " was set on fire"
		bad_events[20] = " was decapitated, temporarily.."
		bad_events[21] = " was tipped by a cow"
		bad_events[22] = " was bucked from a horse"
		bad_events[23] = " was bitten by a møøse"
		bad_events[24] = " was sat on by a giant"
		bad_events[25] = " ate a plate of discounted, day-old sushi"
		bad_events[26] = " got harassed by peer"
		bad_events[27] = " got lost in the woods"
		bad_events[28] = " misplaced his map"
		bad_events[29] = " broke his compass"
		bad_events[30] = " lost his glasses"
		bad_events[31] = " walked face-first into a tree"
		//OpenStack Related
		bad_events[32] = " uploaded a review with a bunch of PRINT statements"
		bad_events[33] = " realised the code he was writing for the last five hours was already in Mitaka"
		bad_events[34] = " walked face-first into a tree"
		bad_events[35] = " walked face-first into a tree"

		  items :=[10]string{"weapon","tunic","shield","leggins","ring","gloves","boots","helm","charm","amulet"}

	*/
	for i := range g.heroes {
		if !g.heroes[i].Enabled {
			continue
		}

		if(rand.Int31n(2000) == 1){

			if(rand.Intn(10) == 1) {  //Ultra Godsend


			} else {


			}
			g.heroes[i].Xpos = truncateInt(g.heroes[i].Xpos+(rand.Intn(3)-1), xMin, xMax)
			g.heroes[i].Ypos = truncateInt(g.heroes[i].Ypos+(rand.Intn(3)-1), yMin, yMax)

		}

	}
}

//Check Hero Level function checks the Hero level and promotes the level is hi/her has reached that level
//Parameters: SQL DB connection Object
//Return: None
func (g *Game) CheckHeroLevel() {

	log.Info("[Engine] CheckHeroLevel Called.................................................")

	sqldb, err := sql.Open("mysql", g.dbconnection)

	if err != nil {
		log.Errorln("[Engine] : CheckHeroLevel : Database settings failed: %s", err)
	}
	defer sqldb.Close()


	for i := range g.heroes {
		if !g.heroes[i].Enabled {
			continue
		}

		if g.heroes[i].NextLevelAt.Before(time.Now()) {
			level := g.heroes[i].Level + 1
			ttl := getTTL(level + 1)
			g.heroes[i].NextLevelAt = time.Now().Add(ttl * time.Second)
			g.heroes[i].Level = level

			var message = g.heroes[i].HeroName + ", " + g.heroes[i].Title + ", has attained Level " +
				strconv.Itoa(level) + "! Next level in " + ttl.String() + " seconds."

			//TODO: Fix the  ttl.String() value

			log.Info(message)

			//Add Event to WorldEvents and HeroWorldEvents Tables
			Insert_World_Event_for_Hero(g.heroes[i].HeroID, message, sqldb)

			//Find a new Item for the Hero
			Find_item( g.heroes[i].HeroID, g.heroes[i].Level, g.heroes[i].HeroName, sqldb)

		}
	}
}

