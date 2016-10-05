package main

import (
  "time"

  log "github.com/Sirupsen/logrus"
)

type Game struct {
  startedAt  time.Time
  heros      []Hero
  adminToken string
  exitChan   chan []byte
}

type Hero struct {
  name      string
  email     string
  class     string
  level     int
  ttl       int
  createdAt time.Time
  equipment Equipment
  xPos      float32
  yPos      float32
}

type Equipment struct {
  ring     int
  amulet   int
  charm    int
  weapon   int
  helm     int
  tunic    int
  gloves   int
  shield   int
  leggings int
  boots    int
}

// NewGame creates a new game
func NewGame() *Game {
  game := &Game{
    startedAt: time.Now(),
    heros:     []Hero{},
    exitChan:  make(chan []byte),
  }
  return game
}

// StartGame starts the game
func StartGame() {
  game := NewGame()

  go game.StartEngine()
  game.StartAPI()
}

// StartEngine starts the engine
func (g *Game) StartEngine() {
  ticker := time.NewTicker(time.Second * 2)
  for {
    select {
    case <-ticker.C:
      log.Info("Ticker")
      g.movePlayers()
    case <-g.exitChan:
      log.Info("Exiting game")
      return
    }
  }
}

func (g *Game) movePlayers() {
  log.Info("Move Players")
}

func (g *Game) authorizeAdmin(token string) bool {
  return g.adminToken == token
}
