package main

import (
  "math/rand"

  log "github.com/Sirupsen/logrus"
)

// GodSend function implements the Gods gits to Heros  on the Realm. It happens 1 an hour and it has 1/4000 chances
// to strike a Hero. The outcome has 80 chances to be good and 20 chances to the bad
func (g *Game) GodSend(trigger int) {

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

  var message string

  for i := range g.heroes {
    if !g.heroes[i].Enabled {
      continue
    }

    if rand.Intn(trigger) == 1 {

      if rand.Intn(10) < 2 { //Ultra Godsend 20%

        //Select a Good Events Random text + Removes time to level up
        message = " WOWO!!!"

      } else { // Upgrade a Weapon

        items := [6]string{"weapon", "tunic", "shield", "leggings", "amulet", "charm"}

        var item_type = rand.Intn(10)
        var item_name = items[item_type]

        switch item_name {

        case "weapon":
          message = " WOWO!!!"
          return

        case "tunic":
          message = " WOWO!!!"
          return

        case "shield":
          message = " WOWO!!!"
          return

        case "leggings":
          message = " WOWO!!!"
          return

        case "amulet":
          message = " WOWO!!!"
          return

        case "charm":
          message = " WOWO!!!"
          return

        default:

        }

      }
    }

    //Add Event to WorldEvents and HeroWorldEvents Tables
    log.Info(message)
    // Insert_World_Event_for_Hero(g.heroes[i].HeroID, message, sqldb)

  }

}
