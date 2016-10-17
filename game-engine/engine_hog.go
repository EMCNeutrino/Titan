package main

import (
  "math/rand"
  "strconv"

  log "github.com/Sirupsen/logrus"
)

// HandOfGod function implements the Gods powers on the Realm. It happens 1 an hour and it has 1/4000 chances
// to strike a Hero. The outcome has 80 chances to be good and 20 chances to the bad
func (g *Game) HandOfGod(trigger int) {

  for i := range g.heroes {
    if !g.heroes[i].Enabled {
      continue
    }

    if rand.Intn(trigger) == 1 {

      var hero_chance = rand.Intn(10)
      var message string
      var hero_next_level = g.heroes[i].Level + 1
      var time_calculation = int(float32((rand.Intn(71) + 5)) / 100 * float32(hero_next_level*3600))

      log.Infof("[Engine] : Hand of Good : time_calculation: %d", time_calculation)

      if hero_chance >= 3 {
        // Good outcome

        message = "Verily I say undo thee, the Heavens have burst forth, and the blessed Hand Of God carried " +
          g.heroes[i].HeroName + " for " + strconv.Itoa(time_calculation) +
          " seconds toward level " + strconv.Itoa(hero_next_level)

        //TODO: Update the Hero TTL for the next level

      } else { //Bad outcome

        message = "Thereupon He stretched out His little finger among them and consummed " +
          g.heroes[i].HeroName + " with fier, slowing the heathen " + strconv.Itoa(time_calculation) +
          " seconds from  level " + strconv.Itoa(hero_next_level)

        //TODO: Update the Hero TTL for the next level
      }

      log.Info("[Hand of God]" + message)

      //Add Event to WorldEvents and HeroWorldEvents Tables
      // Insert_World_Event_for_Hero(g.heroes[i].HeroID, message, sqldb)

    }

  }
}
