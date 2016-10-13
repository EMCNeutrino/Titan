package main

import (
  "fmt"
  "math"
  "math/rand"
  "time"

  log "github.com/Sirupsen/logrus"
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
  heroes           []Hero
  adminToken       string
  joinChan         chan JoinRequest
  activateHeroChan chan ActivateHeroRequest
  exitChan         chan []byte
}

type Hero struct {
  Name        string `json:"name"`
  Email       string `json:"email"`
  Class       string `json:"class"`
  Enabled     bool   `json:"enabled"`
  token       string
  Level       int `json:"level"`
  nextLevelAt time.Time
  createdAt   time.Time
  Equipment   Equipment `json:"equipment"`
  Xpos        int       `json:"x_pos"`
  Ypos        int       `json:"y_pos"`
}

type Equipment struct {
  Ring     int
  Amulet   int
  Charm    int
  Weapon   int
  Helm     int
  Tunic    int
  Gloves   int
  Shield   int
  Leggings int
  Boots    int
}

// NewGame creates a new game
func NewGame(adminToken string) *Game {
  game := &Game{
    startedAt:        time.Now(),
    heroes:           []Hero{},
    joinChan:         make(chan JoinRequest),
    activateHeroChan: make(chan ActivateHeroRequest),
    exitChan:         make(chan []byte),
    adminToken:       adminToken,
  }
  return game
}

// StartGame starts the game
func StartGame(adminToken string) {
  game := NewGame(adminToken)

  go game.StartEngine()
  game.StartAPI()
}

// StartEngine starts the engine
func (g *Game) StartEngine() {
  ticker := time.NewTicker(time.Second * 2)
  for {
    select {
    case <-ticker.C:
      g.moveHeroes()
      g.checkLevels()
      //TODO: check battles
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

  hero := &Hero{
    Name:        name,
    Email:       email,
    Class:       class,
    Enabled:     false,
    token:       randToken(),
    Level:       0,
    nextLevelAt: time.Now().Add(99999 * time.Hour),
    createdAt:   time.Now(),
    Equipment: Equipment{
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
    },
    Xpos: rand.Intn(xMax-xMin) + xMin,
    Ypos: rand.Intn(yMax-yMin) + yMin,
  }

  g.heroes = append(g.heroes, *hero)
  log.Infof("Hero %s has been created, but will not play until it's activated.", hero.Name)
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

  ttl := getTTL(1) // Time to level 1
  g.heroes[i].Enabled = true
  g.heroes[i].nextLevelAt = time.Now().Add(ttl * time.Second)
  log.Infof("Success! Hero %s has been activated and will reach level 1 in %d seconds.", g.heroes[i].Name, ttl)
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

func (g *Game) checkLevels() {
  for i := range g.heroes {
    if !g.heroes[i].Enabled {
      continue
    }

    if g.heroes[i].nextLevelAt.Before(time.Now()) {
      level := g.heroes[i].Level + 1
      ttl := getTTL(level + 1)
      g.heroes[i].nextLevelAt = time.Now().Add(ttl * time.Second)
      g.heroes[i].Level = level
      log.Infof("Hero %s reached level %d. Next level in %d seconds.", g.heroes[i].Name, level, ttl)
    }
  }
}

func (g *Game) authorizeAdmin(token string) bool {
  return g.adminToken == token
}

func (g *Game) getHeroIndex(name string) int {
  for i, hero := range g.heroes {
    if hero.Name == name {
      return i
    }
  }
  return -1
}

func getTTL(level int) time.Duration {
  return time.Duration(levelUpSeconds * (math.Pow(levelUpBase, float64(level))))
}
