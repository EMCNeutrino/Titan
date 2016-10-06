package main

import (
  "fmt"
  "math/rand"
  "time"

  log "github.com/Sirupsen/logrus"
)

const (
  xMax = 500
  yMax = 500
  xMin = 0
  yMin = 0
)

type Game struct {
  startedAt  time.Time
  heros      []Hero
  adminToken string
  joinChan   chan string
  exitChan   chan []byte
}

type Hero struct {
  Name      string `json:"name"`
  Email     string `json:"email"`
  Class     string `json:"class"`
  enabled   bool
  token     string
  Level     int `json:"level"`
  ttl       int
  createdAt time.Time
  Equipment Equipment `json:"equipment"`
  Xpos      int       `json:"x_pos"`
  Ypos      int       `json:"y_pos"`
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
    startedAt:  time.Now(),
    heros:      []Hero{},
    joinChan:   make(chan string),
    exitChan:   make(chan []byte),
    adminToken: adminToken,
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
      g.movePlayers()
    case token := <-g.joinChan:
      log.Info("Join hero")
      if err := g.joinPlayer("asd", "asdads@asd.com", "asdd class", token); err != nil {
        log.Error(err)
      }
    case <-g.exitChan:
      log.Info("Exiting game")
      return
    }
  }
}

func (g *Game) joinPlayer(name, email, class, adminToken string) error {

  if !g.authorizeAdmin(adminToken) {
    return fmt.Errorf("You are not authorized to perform this action.")
  }

  hero := &Hero{
    Name:      name,
    Email:     email,
    Class:     class,
    enabled:   false,
    token:     randToken(),
    Level:     1,
    ttl:       1,
    createdAt: time.Now(),
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
    Xpos: 0,
    Ypos: 0,
  }

  g.heros = append(g.heros, *hero)
  return nil
}

func (g *Game) activatePlayer() {

}

func (g *Game) movePlayers() {
  for i := range g.heros {
    g.heros[i].Xpos = truncateInt(g.heros[i].Xpos+(rand.Intn(3)-1), xMin, xMax)
    g.heros[i].Ypos = truncateInt(g.heros[i].Ypos+(rand.Intn(3)-1), yMin, yMax)
  }
}

func (g *Game) authorizeAdmin(token string) bool {
  return g.adminToken == token
}
