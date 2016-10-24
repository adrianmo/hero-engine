package main

import (
  "fmt"
  "math/rand"
)

func (g *Game) calamity() {

  var calamity = []string{

    " was bitten by Neutron",
    " fell into a hole",
    " bit their tongue",
    " set thyself on fire",
    " ate a poisonous fruit",
    " lost their mind",
    " died, temporarily..",
    " was caught in a terrible snowstorm",
    " EXPLODED, somewhat..",
    " got knifed in a dark alley",
    " saw an episode of Ally McBeal",
    " got turned INSIDE OUT, practically",
    " ate a very disagreeable fruit, getting a terrible case of heartburn",
    " met up with a mob hitman for not paying his hosting bills",
    " has fallen ill with the black plague",
    " was struck by lightning",
    " was attacked by a rabid giant rabbit",
    " was attacked by a rabid wolverine",
    " was set on fire",
    " was decapitated, temporarily..",
    " was tipped by a cow",
    " was bucked from a horse",
    " was bitten by a møøse",
    " was sat on by a giant",
    " ate a plate of discounted, day-old sushi",
    " got harassed by peer",
    " got lost in the woods",
    " misplaced his map",
    " broke his/her compass",
    " lost his/her glasses",
    " walked face-first into a tree",
    " uploaded a review with a bunch of PRINT statements",
    " realised the code he was writing for the last five hours was already in Mitaka",
    " walked face-first into a tree",
  }

  for i := range g.heroes {

    if !g.heroes[i].Enabled {
      continue
    }

    if rand.Intn(500) == 0 {

      var message string

      if rand.Intn(10) < 2 { //Ultra Calamity 20%

        //Select a Calamity Events Random text + Removes time to level up
        calamityEventID := rand.Intn(len(calamity))
        seconds := rand.Intn(50) * (g.heroes[i].Level + 1)
        g.heroes[i].updateTTL(0 - seconds)

        message = fmt.Sprintf("%s %s. This terrible calamity has slowed him/her %d seconds for the next Level!", g.heroes[i].HeroName, calamity[calamityEventID], seconds)

      } else { // Upgrade a Weapon

        items := [6]string{"weapon", "tunic", "shield", "leggings", "amulet", "charm"}
        var itemType = items[rand.Intn(6)]

        g.heroes[i].updateItem(itemType, int(float64(g.heroes[i].Equipment.Weapon)*0.9))

        switch itemType {
        case "weapon":
          message = fmt.Sprintf("%s left his weapon out in the rain to rust! %s's weapon loses 10%% effectiveness.", g.heroes[i].HeroName, g.heroes[i].HeroName)
        case "tunic":
          message = fmt.Sprintf("%s spilled a level 7 shrinking potion on his tunic! %s's tunic loses 10%% effectiveness.", g.heroes[i].HeroName, g.heroes[i].HeroName)
        case "shield":
          message = fmt.Sprintf("%s's shield was damaged by a dragon's fiery breath! %s's shield loses 10%% effectiveness.", g.heroes[i].HeroName, g.heroes[i].HeroName)
        case "leggings":
          message = fmt.Sprintf("%s' burned a hole through his leggings while ironing them! %s's leggings loses 10%% effectiveness.", g.heroes[i].HeroName, g.heroes[i].HeroName)
        case "amulet":
          message = fmt.Sprintf("%s fell, chipping the stone in his amulet! %s's amulet loses 10%% effectiveness.", g.heroes[i].HeroName, g.heroes[i].HeroName)
        case "charm":
          message = fmt.Sprintf("%s slipped and dropped his charm in a dirty bog! %s's charm loses 10%% effectiveness.", g.heroes[i].HeroName, g.heroes[i].HeroName)
        }
      }

      if len(message) > 0 {
        g.sendEvent(message, g.heroes[i])
      }
    }
  }
}
