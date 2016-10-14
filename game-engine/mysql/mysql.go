
//Package provides the Data Access Layer (DAL) for the Hero Game.
package main

import (

	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"encoding/json"
	"os"
	"fmt"
	"strconv"
	log "github.com/Sirupsen/logrus"
	"math/rand"
)

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
func Db_connection() string {
// Returns the Connection String for the Database

	file,err := os.Open("dbconfig.json")

	if err != nil {
		fmt.Println("File Error:", err)
	}

	decoder := json.NewDecoder(file)

	var db_conf = Db_configuration{}

	err2 := decoder.Decode(&db_conf)

	if err2 != nil {
		fmt.Println("Decoder Error:", err2)
	}

	return db_conf.Db_user + ":" + db_conf.Db_pass + "@tcp(" + db_conf.Db_ip + ":" + db_conf.Db_port + ")/" + db_conf.Db_name

}



func main() {

	var dbconn = Db_connection()

	log.Info("Database Connection: " + dbconn)

	sqldb, err := sql.Open("mysql", dbconn)

	if err != nil {
		log.Errorln("Database settings failed: %s", err)
	}

	defer sqldb.Close()

	// Open doesn't open a connection. Validate DSN data:
	err = sqldb.Ping()

	if err != nil {
		log.Errorln("Database Ping failed: %s", err)
	}

	//Find an Item for Hero 1
	for i := 0; i < 100; i++ {
		log.Info("Hero # " + strconv.Itoa(i))
		Find_item(rand.Intn(100)+1, rand.Intn(70), "Gandalf", sqldb)
	}


}


// Find_item generates a new item for the hero when they level up
// and notifies the player about the item found
func Find_item(heroID int, hero_level int, hero_name string, conn *sql.DB){

	//log.Info("Find Items called, Hero Level: " + strconv.Itoa(hero_level))

	items :=[10]string{"weapon","tunic","shield","leggins","ring","gloves","boots","helm","charm","amulet"}

	find_chance := [51]float32{ 100.00, 91.93227152,84.51542547,77.69695042,71.42857143,65.66590823,60.36816105,55.49782173,
		51.02040816,46.90422016,43.12011504,39.64130124,36.44314869,33.5030144,	30.80008217,28.31521517,26.03082049,
		23.93072457,22.00005869,20.22515369,18.59344321,17.0933747,	15.71432764,14.44653835,13.28103086,12.20955335,
		11.22451974,10.31895597,9.486450616,8.721109539,8.017514101,7.370682832,6.776036155,6.229363956,5.726795786,
		5.264773452,4.840025825,4.449545683,4.090568419,3.760552466,3.457161303,3.178246916,2.921834585,2.686108904,
		2.469400931,2.270176369,2.087024703,1.918649217,1.763857808,1.621554549,1.490731931	}

	var item_type int
	var item_name string
	var item_level int
	var item_found_chance float32

	for i := hero_level; i > 0; i-- {

		if(i > 50) {
			//After Hero Level of 50, has a 1% chance to find an item.
			item_found_chance = 1.0

		} else {

			item_found_chance = find_chance[i]
		}
		//Start with highest Level Item and subtract a level as it misses the chance
		if(rand.Intn(100) <= int(item_found_chance)){

			item_gain_percentage := float64(rand.Intn(100))
			item_level = int(float64(i) + (float64(i) * (item_gain_percentage/100)))
			item_type = rand.Intn(10)
			item_name = items[item_type]

			var log_msg = "Item Found: " + item_name + " | Hero Level: " + strconv.Itoa(hero_level)+ " found @ Level: " + strconv.Itoa(i)+ " Item Level: " + strconv.Itoa(item_level)

			log.Info(log_msg)

			break
		}
	}

	//check if No items where found

	var current_item_level = Get_Item_By_HeroID(heroID, item_name, conn)

	log.Info("Items: Current: " + strconv.Itoa(current_item_level) + " | New: " + strconv.Itoa(item_level))

	var message string
	var message_plural string

	if (item_name == items[3] || item_name == items[5] || item_name == items[6]) {

		message_plural =  " are only level "

	} else {

		message_plural =  " is only level "
	}

	if item_level > current_item_level {

		//Replace the current item value with the new one
		Update_Item_for_Hero(heroID,item_name, item_level, conn)

		message = "You found a level " + strconv.Itoa(item_level) + " " + item_name + "! Your current " + item_name + message_plural +  strconv.Itoa(current_item_level) + ", so it seems Luck is with you!"

	} else {

		//Message back to player that current item level is better
		message = "You found a level " + strconv.Itoa(item_level) + " " + item_name + "! Your current " + item_name + message_plural +  strconv.Itoa(current_item_level) + ", so it seems Luck is against you. You toss the " + item_name + "."

	}

	log.Info(message)

	//Message back to player that new item is better
	Insert_World_Event_for_Hero(heroID, message, conn)


}

func Get_Item_By_HeroID(heroID int, item_type string, conn *sql.DB) int {

	log.Info("Get_item_level_for_user ")

	var current_item_level int

	//Check what is the level of the current Item, update value if needed, msg player
	var query = "SELECT " + item_type + " FROM item WHERE hero_id = ?"

	log.Info("Select Query: " + query)

	stmt, err := conn.Prepare(query)
	if err != nil {
		log.Errorln("DB: Get_Item_By_HeroID: Prepare Query failed: %s", err)
	}

	err = stmt.QueryRow(heroID).Scan(&current_item_level)

	if err != nil {
		log.Errorln("DB: Get_Item_By_HeroID: QueryRow failed: %s", err)
	}

	log.Info("Item Value: " + strconv.Itoa(current_item_level))

	defer stmt.Close()

	return current_item_level
}

func Update_Item_for_Hero(heroID int, item_type string, item_level int, conn *sql.DB) {

	log.Info("Updating Item: " + item_type + " with Level: " + strconv.Itoa(item_level) +" from Hero ID: " + strconv.Itoa(heroID))

	var query = "UPDATE item SET " + item_type + "=" + strconv.Itoa(item_level) +" WHERE hero_id="+ strconv.Itoa(heroID)

	_, err := conn.Exec(query)

	if err != nil {
		log.Errorln("DB: Insert_Item_for_Hero: Item Insert failed: %s", err)
	}


	log.Info("Executed: " + query)
}

func Insert_World_Event_for_Hero(heroID int, worldEvent string, conn *sql.DB) {

	log.Info("Inserting World Event: HeroID: " + strconv.Itoa(heroID) + " | Event: " + worldEvent)

	var query = "INSERT INTO worldevent (event_text) VALUES ('" + worldEvent + "')"

	log.Info("World Event Query: " + query)

	statement, err := conn.Exec(query)

	last_worldevent_Id, err := statement.LastInsertId()

	if err != nil {
		log.Errorln("DB: Insert_Item_for_Hero: Worldevent Insert failed: %s", err)
	}

	if last_worldevent_Id != 0 {

		Insert_Hero_World_Event(heroID, int(last_worldevent_Id), conn)
	}
}

func Insert_Hero_World_Event(heroID int, woldEvent_id int, conn *sql.DB) {

	var query = "INSERT INTO heroworldevent (hero_id, worldevent_id ) VALUES (" + strconv.Itoa(heroID) + ", " + strconv.Itoa(woldEvent_id) + ")"

	log.Info("Heroworldevent Query: " + query)

	_, err := conn.Exec(query)

	if err != nil {
		log.Errorln("DB: Insert_Item_for_Hero: Heroworldevent Insert failed: %s", err)
	}
}







