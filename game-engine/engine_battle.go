package main

import (
  "fmt"
  "math"
  "math/rand"
  "time"

  log "github.com/Sirupsen/logrus"
)

const (
  battleCooldown = time.Duration(1) * time.Minute
  battleDistance = 5
)

// battle function implements the battle logic
func (g *Game) checkBattles() {

  // Shuffle heroes list
  heroesShuffle := make([]*Hero, len(g.heroes))
  copy(heroesShuffle, g.heroes)
  perm := rand.Perm(len(g.heroes))
  for i, v := range perm {
    heroesShuffle[v] = g.heroes[i]
  }

  for _, h1 := range heroesShuffle {

    if !h1.Enabled {
      continue
    }

    if h1.lastBattleAt.Add(battleCooldown).After(time.Now()) {
      // Hero fought very recently
      log.Debugf("[Battle] %s fought very recently. Skipping", h1.HeroName)
      continue
    }

    for _, h2 := range g.heroes {
      if h1 == h2 {
        continue
      }

      if !h2.Enabled {
        continue
      }

      if h2.lastBattleAt.Add(battleCooldown).After(time.Now()) {
        // Hero fought very recently
        log.Debugf("[Battle] %s fought very recently. Skipping", h2.HeroName)
        continue
      }

      if heroesDistance(h1, h2) > battleDistance {
        // Too far away
        log.Debugf("[Battle] %s and %s are too far away.", h1.HeroName, h2.HeroName)
        continue
      }

      //TODO: battle logic

      // H1 lost
      h1.lastBattleAt = time.Now()
      h2.lastBattleAt = time.Now()
      h1.updateTTL(1000)
      h2.updateTTL(-1000)
      message := fmt.Sprintf("%s fought to %s and lost. Incresing the time to the next level by %d seconds.", h1.HeroName, h2.HeroName, 1000)
      go g.sendEvent(message, h1, h2)
    }
  }
}

func battle(h1 *Hero, h2 *Hero) {

}

func heroesDistance(h1 *Hero, h2 *Hero) float64 {
  xPow := math.Pow(float64(h1.Xpos-h2.Xpos), 2)
  yPow := math.Pow(float64(h1.Ypos-h2.Ypos), 2)
  distance := math.Sqrt(xPow + yPow)
  log.Debugf("[Battle] Hero 1: %s (%d,%d) | Hero 2: %s (%d,%d) | Distance: %.2f", h1.HeroName, h1.Xpos, h1.Ypos, h2.HeroName, h2.Xpos, h2.Ypos, distance)
  return distance
}
