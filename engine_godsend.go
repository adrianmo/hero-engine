package main

import (
  "math/rand"

  log "github.com/Sirupsen/logrus"
  "fmt"

)

// GodSend function implements the Gods gits to Heros  on the Realm. It happens 1 an hour and it has 1/4000 chances
// to strike a Hero. The outcome has 80 chances to be good and 20 chances to the bad
func (g *Game) GodSend() {

   log.Debug("[GodSend] Calculate Godsend")

  var good_events = map[int]string{}

  good_events[1] = " found a pair of nice Shoes"
  good_events[2] = " caught a Unicorn"
  good_events[3] = " discovered a secret, underground passage leading to Barcelona's best tavern"
  good_events[4] = " was taught to run quickly by a secret tribe of pygmies that know how to, among other things, run quickly"
  good_events[5] = " discovered caffeinated coffee"
  good_events[6] = " grew an extra leg"
  good_events[7] = " was visited by a very pretty nymph"
  good_events[8] = " found pretty kitten"
  good_events[9] = " learned Python"
  good_events[10] = " found an exploit in the Neutrino Heros Idle RPG code"
  good_events[11] = " tamed a wild horse"
  good_events[12] = " found a one-time-use spell of quickness"
  good_events[13] = " bought a faster computer"
  good_events[14] = " bribed the local OpenStack administrator"
  good_events[15] = " stopped using dial-up"
  good_events[16] = " invented the wheel"
  good_events[17] = " gained a sixth sense"
  good_events[18] = " got a kiss from an Angel"
  good_events[19] = " had his clothes laundered by a passing fairy"
  good_events[20] = " was rejuvenated by drinking from a magic stream"
  good_events[21] = " was bitten by a radioactive spider"
  good_events[22] = " was invited to dance a Sardana by Barcelona's Cathedral"
  //OpenSTack Related
  good_events[23] = " was accepted into the Leage of fantastic Stakers"
  good_events[24] = " was notified that Jenkins tests passed"
  good_events[25] = " got his first patch +4 approved in OpenStack"
  good_events[26] = " got a HEAT template successfully deployed"

  var message string

  for i := range g.heroes {

    if !g.heroes[i].Enabled {
      continue
    }

    if rand.Intn(1) == 1 {

      if rand.Intn(10) < 2 { //Ultra Godsend 20%

        //Select a Good Events Random text + Removes time to level up
        var good_eventID = rand.Intn(26)

        var ultra_gain = int(float32((rand.Intn(8) + 5)) / 100 * float32(g.heroes[i].Level+1))

        message = fmt.Sprintf("%s,%s. This wondrous godsend has accelerated them %d", g.heroes[i].HeroName, good_events[good_eventID], ultra_gain)
        return

      } else { // Upgrade a Weapon

        items := [6]string{"weapon", "tunic", "shield", "leggings", "amulet", "charm"}

        var item_type = rand.Intn(6)
        var item_name = items[item_type]

        switch item_name {

          case "weapon":
            message = fmt.Sprintf("%s sharpened the edge of his weapon! %s's weapon gains 10% effectiveness.", g.heroes[i].HeroName, g.heroes[i].HeroName)
            g.heroes[i].Equipment.Weapon = int(float64(g.heroes[i].Equipment.Weapon) * 1.1)
            return

          case "tunic":
            message = fmt.Sprintf("A magician cast a spell of Rigidity on %s's tunic! %s tunic gains 10% effectiveness.", g.heroes[i].HeroName, g.heroes[i].HeroName)
            g.heroes[i].Equipment.Weapon = int(float64(g.heroes[i].Equipment.Tunic) * 1.1)
            return

          case "shield":
            message = fmt.Sprintf("%s reinforced his shield with a dragon's scales! %s shield gains 10% effectiveness.", g.heroes[i].HeroName, g.heroes[i].HeroName)
            g.heroes[i].Equipment.Weapon = int(float64(g.heroes[i].Equipment.Shield) * 1.1)
            return

          case "leggings":
            message = fmt.Sprintf("The local wizard imbued %s's pants with a Spirit of Fortitude! %s leggings gain 10% effectiveness.", g.heroes[i].HeroName, g.heroes[i].HeroName)
            g.heroes[i].Equipment.Weapon = int(float64(g.heroes[i].Equipment.Leggings) * 1.1)
            return

          case "amulet":
            message = fmt.Sprintf("%s's amulet was blessed by a passing cleric! %s's amulet gains 10% effectiveness.", g.heroes[i].HeroName, g.heroes[i].HeroName)
            g.heroes[i].Equipment.Weapon = int(float64(g.heroes[i].Equipment.Amulet) * 1.1)
            return

          case "charm":
            message = fmt.Sprintf("%s's charm was enchanted by the Queen of fairies! %s's charm gains 10% effectiveness.", g.heroes[i].HeroName, g.heroes[i].HeroName)
            g.heroes[i].Equipment.Weapon = int(float64(g.heroes[i].Equipment.Charm) * 1.1)
            return

        }

      }
    }

    //Add Event to WorldEvents and HeroWorldEvents Tables
    log.Info(message)
    // Insert_World_Event_for_Hero(g.heroes[i].HeroID, message, sqldb)

  }

}
