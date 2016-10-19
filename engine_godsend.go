package main

import (
  "math/rand"
  "fmt"
)

// GodSend function implements the Gods gits to Heros  on the Realm. It happens 1 an hour and it has 1/4000 chances
// to strike a Hero. The outcome has 80 chances to be good and 20 chances to the bad
func (g *Game) GodSend() {

  var goodEvents = []string{
    "found a pair of nice Shoes",
    "caught a Unicorn",
    "discovered a secret, underground passage leading to Barcelona's best tavern",
    "was taught to run quickly by a secret tribe of pygmies that know how to, among other things, run quickly",
    "discovered caffeinated coffee",
    "grew an extra leg",
    "was visited by a very pretty nymph",
    "found pretty kitten",
    "learned Python",
    "found an exploit in the Neutrino Heros Idle RPG code",
    "tamed a wild horse",
    "found a one-time-use spell of quickness",
    "bought a faster computer",
    "bribed the local OpenStack administrator",
    "stopped using dial-up",
    "invented the wheel",
    "gained a sixth sense",
    "got a kiss from an Angel",
    "had his clothes laundered by a passing fairy",
    "was rejuvenated by drinking from a magic stream",
    "was bitten by a radioactive spider",
    "was invited to dance a Sardana by Barcelona's Cathedral",
    "was accepted into the Leage of fantastic Stakers",
    "was notified that Jenkins tests passed",
    "got his first patch +4 approved in OpenStack",
    "got a HEAT template successfully deployed",
  }

  for i := range g.heroes {

    if !g.heroes[i].Enabled {
      continue
    }

    if rand.Intn(1) == 0 {

      var message string

      if rand.Intn(10) < 2 { //Ultra Godsend 20%

        //Select a Good Events Random text + Removes time to level up
        goodEventID := rand.Intn(len(goodEvents))
        seconds := rand.Intn(50) * (g.heroes[i].Level + 1)
        g.heroes[i].updateTTL(0 - seconds)

        message = fmt.Sprintf("%s %s. This wondrous godsend has accelerated him/her %d seconds for the next Level!", g.heroes[i].HeroName, goodEvents[goodEventID], seconds)

      } else { // Upgrade a Weapon

        items := [6]string{"weapon", "tunic", "shield", "leggings", "amulet", "charm"}
        var itemType = items[rand.Intn(6)]

        g.heroes[i].updateItem(itemType, int(float64(g.heroes[i].Equipment.Weapon)*1.1))

        switch itemType {
        case "weapon":
          message = fmt.Sprintf("%s sharpened the edge of his weapon! %s's weapon gains 10%% effectiveness.", g.heroes[i].HeroName, g.heroes[i].HeroName)
        case "tunic":
          message = fmt.Sprintf("A magician cast a spell of Rigidity on %s's tunic! %s's tunic gains 10%% effectiveness.", g.heroes[i].HeroName, g.heroes[i].HeroName)
        case "shield":
          message = fmt.Sprintf("%s reinforced his shield with a dragon's scales! %s's shield gains 10%% effectiveness.", g.heroes[i].HeroName, g.heroes[i].HeroName)
        case "leggings":
          message = fmt.Sprintf("The local wizard imbued %s's pants with a Spirit of Fortitude! %s's leggings gain 10%% effectiveness.", g.heroes[i].HeroName, g.heroes[i].HeroName)
        case "amulet":
          message = fmt.Sprintf("%s's amulet was blessed by a passing cleric! %s's amulet gains 10%% effectiveness.", g.heroes[i].HeroName, g.heroes[i].HeroName)
        case "charm":
          message = fmt.Sprintf("%s's charm was enchanted by the Queen of fairies! %s's charm gains 10%% effectiveness.", g.heroes[i].HeroName, g.heroes[i].HeroName)
        }
      }

      if len(message) > 0 {
        g.sendEvent("[Godsend] "+message, g.heroes[i])
      }
    }
  }
}
