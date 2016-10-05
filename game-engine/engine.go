package main

import (
  "time"

  log "github.com/Sirupsen/logrus"
)

type Engine struct {
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

// NewEngine creates a new game engine
func NewEngine() *Engine {
  engine := &Engine{
    startedAt: time.Now(),
    heros:     []Hero{},
    exitChan:  make(chan []byte),
  }
  return engine
}

// StartEngine starts the game engine
func StartEngine() {
  engine := NewEngine()

  go engine.Start()
}

// Start starts the engine
func (e *Engine) Start() {
  ticker := time.NewTicker(time.Second * 2)
  for {
    select {
    case <-ticker.C:
      log.Info("Ticker")
      e.movePlayers()
    case <-e.exitChan:
      log.Info("Exiting game")
      return
    }
  }
}

func (e *Engine) movePlayers() {
  log.Info("Move Players")
}

func (e *Engine) authorizeAdmin(token string) bool {
  return e.adminToken == token
}
