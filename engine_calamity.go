package main

import "math/rand"

func (g *Game) Calamity() {

  //TODO Complete Calamity

  /*
     var bad_events = map[int]string{}

     bad_events[1] = " was bitten by Neutron"
     bad_events[2] = " fell into a hole"
     bad_events[3] = " bit their tongue"
     bad_events[4] = " set thyself on fire"
     bad_events[5] = " ate a poisonous fruit"
     bad_events[6] = " lost their mind"
     bad_events[7] = " died, temporarily.."
     bad_events[8] = " was caught in a terrible snowstorm"
     bad_events[9] = " EXPLODED, somewhat.."
     bad_events[10] = " got knifed in a dark alley"
     bad_events[11] = " saw an episode of Ally McBeal"
     bad_events[12] = " got turned INSIDE OUT, practically"
     bad_events[13] = " ate a very disagreeable fruit, getting a terrible case of heartburn"
     bad_events[14] = " met up with a mob hitman for not paying his hosting bills"
     bad_events[15] = " has fallen ill with the black plague"
     bad_events[16] = " was struck by lightning"
     bad_events[17] = " was attacked by a rabid giant rabbit"
     bad_events[18] = " was attacked by a rabid wolverine"
     bad_events[19] = " was set on fire"
     bad_events[20] = " was decapitated, temporarily.."
     bad_events[21] = " was tipped by a cow"
     bad_events[22] = " was bucked from a horse"
     bad_events[23] = " was bitten by a møøse"
     bad_events[24] = " was sat on by a giant"
     bad_events[25] = " ate a plate of discounted, day-old sushi"
     bad_events[26] = " got harassed by peer"
     bad_events[27] = " got lost in the woods"
     bad_events[28] = " misplaced his map"
     bad_events[29] = " broke his compass"
     bad_events[30] = " lost his glasses"
     bad_events[31] = " walked face-first into a tree"
     //OpenStack Related
     bad_events[32] = " uploaded a review with a bunch of PRINT statements"
     bad_events[33] = " realised the code he was writing for the last five hours was already in Mitaka"
     bad_events[34] = " walked face-first into a tree"
     bad_events[35] = " walked face-first into a tree"

       items :=[10]string{"weapon","tunic","shield","leggins","ring","gloves","boots","helm","charm","amulet"}

  */
  for i := range g.heroes {
    if !g.heroes[i].Enabled {
      continue
    }

    if rand.Int31n(2000) == 1 {

      if rand.Intn(10) == 1 { //Ultra Godsend

      } else {

      }
      g.heroes[i].Xpos = truncateInt(g.heroes[i].Xpos+(rand.Intn(3)-1), xMin, xMax)
      g.heroes[i].Ypos = truncateInt(g.heroes[i].Ypos+(rand.Intn(3)-1), yMin, yMax)

    }
  }
}
