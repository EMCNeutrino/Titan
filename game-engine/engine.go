package main

import (
  "fmt"
  "math/rand"
  "os"
  "time"

  log "github.com/Sirupsen/logrus"
)

const (
  xMax           = 500
  yMax           = 500
  xMin           = 0
  yMin           = 0
  levelUpSeconds = 10
  levelUpBase    = float64(1.16)
)

type Game struct {
  startedAt        time.Time
  heroes           []Hero
  adminToken       string
  joinChan         chan JoinRequest
  activateHeroChan chan ActivateHeroRequest
  exitChan         chan []byte
  databaseURL      string
}

// NewGame creates a new game
func NewGame(adminToken string) (*Game, error) {

  databaseURL := os.Getenv("DATABASE_URL")
  if databaseURL == "" {
    return nil, fmt.Errorf("Missing environment variable DATABASE_URL")
  }

  game := &Game{
    startedAt:        time.Now(),
    heroes:           []Hero{},
    joinChan:         make(chan JoinRequest),
    activateHeroChan: make(chan ActivateHeroRequest),
    exitChan:         make(chan []byte),
    adminToken:       adminToken,
    databaseURL:      databaseURL,
  }

  return game, nil
}

// StartGame starts the game
func StartGame(adminToken string) {
  game, err := NewGame(adminToken)
  if err != nil {
    log.Panic(err)
  }

  err = LoadFromDB(game)
  if err != nil {
    log.Panic(err)
  }

  go game.StartEngine()
  game.StartAPI()
}

// StartEngine starts the engine
func (g *Game) StartEngine() {

  ticker := time.NewTicker(time.Second * 2)
  tickerDB := time.NewTicker(time.Minute * 1)

  for {
    select {
    case <-ticker.C:
      g.moveHeroes()
      g.checkLevels()
      // g.HandOfGod(4000 * 3600)
    //TODO: check battles
    case <-tickerDB.C:
      log.Info("Saving game state to DB")
      if err := SaveToDB(g); err != nil {
        log.Error(err)
      }
    case req := <-g.joinChan:
      log.Info("Join hero")
      success, message := g.joinHero(req.firstName, req.lastName, req.email, req.twitter, req.heroName, req.heroClass, req.TokenRequest.token)
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

func (g *Game) joinHero(firstName, lastName, email, twitter, heroName, heroClass, adminToken string) (bool, string) {

  if !g.authorizeAdmin(adminToken) {
    return false, "You are not authorized to perform this action."
  }

  hero := &Hero{
    FirstName:   firstName,
    LastName:    lastName,
    Email:       email,
    Twitter:     twitter,
    HeroName:    heroName,
    HeroClass:   heroClass,
    Enabled:     false,
    token:       randToken(),
    Level:       0,
    nextLevelAt: ttlToDatetime(99999 * time.Hour),
    CreatedAt:   time.Now(),
    Equipment: &Equipment{
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
      Total:    0,
    },
    Xpos: rand.Intn(xMax-xMin) + xMin,
    Ypos: rand.Intn(yMax-yMin) + yMin,
  }

  g.heroes = append(g.heroes, *hero)
  log.Infof("Hero %s has been created, but will not play until it's activated.", hero.HeroName)
  return true, fmt.Sprintf("Token: %s", hero.token)
}

func (g *Game) activateHero(name, token string) bool {

  i := g.getHeroIndex(name)
  if i == -1 {
    return false
  }
  if g.heroes[i].token != token {
    return false
  }

  ttl := getTTLForLevel(1) // Time to level 1
  g.heroes[i].nextLevelAt = ttlToDatetime(ttl)
  g.heroes[i].Enabled = true
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

func (g *Game) getHero(name string) (Hero, error) {
  for _, hero := range g.heroes {
    if hero.HeroName == name {
      return hero, nil
    }
  }
  return Hero{}, fmt.Errorf("Hero not found")
}

func (g *Game) sendEvent(message string, heroes ...*Hero) {
  log.Infof("Event: %s", message)

  g.saveEventToDB(message, heroes)
}

/*
//WORLD EVENTS

EVENT			Frequency
Hand of God 	20 hours
Team Battle		24 hours
Calamity 		8 hours
GodSend			4 hours
*/

// checkLevels checks the Hero level and promotes the level is hi/her has reached that level
func (g *Game) checkLevels() {

  for i := range g.heroes {
    if !g.heroes[i].Enabled {
      continue
    }

    if g.heroes[i].nextLevelAt.Before(time.Now()) {
      level := g.heroes[i].Level + 1
      ttl := getTTLForLevel(level + 1)
      g.heroes[i].nextLevelAt = ttlToDatetime(ttl)
      g.heroes[i].Level = level

      message := fmt.Sprintf("%s has attained level %d! Next level in %d seconds.", g.heroes[i].HeroName, level, ttl)

      g.sendEvent(message, &g.heroes[i])
      g.findItem(&g.heroes[i])
    }
  }
}
