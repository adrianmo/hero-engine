package main

import (
  "fmt"
  "math/rand"

  log "github.com/Sirupsen/logrus"
)

// handOfGod function implements the Gods powers on the Realm. It happens 1 an hour and it has 1/4000 chances
// to strike a Hero. The outcome has 80 chances to be good and 20 chances to the bad
func (g *Game) handOfGod() {

  for i := range g.heroes {
    if !g.heroes[i].Enabled {
      continue
    }

    if rand.Intn(10) == 1 {
      var message string
      var timeCalculation = int(float32((rand.Intn(71) + 5)) / 100 * float32((g.heroes[i].Level+1)*3600))

      log.Infof("[Hand of God]: Hero: %s.  timeCalculation: %d", timeCalculation)

      if rand.Intn(10) >= 3 {
        // Good outcome
        message = fmt.Sprintf("Verily I say undo thee, the Heavens have burst forth, and the blessed Hand Of God carried "+
          "%s for %d seconds toward level %d", g.heroes[i].HeroName, timeCalculation, g.heroes[i].Level+1)

        g.heroes[i].updateTTL(0 - timeCalculation)
      } else {
        //Bad outcome
        message = fmt.Sprintf("Thereupon He stretched out his little finger among them and consummed "+
          "%s with fier, slowing the heathen %d"+
          " seconds from level %d", g.heroes[i].HeroName, timeCalculation, g.heroes[i].Level+1)

        g.heroes[i].updateTTL(timeCalculation)
      }

      log.Infof("[Engine_hog] : Hand of Good : TTL after: %d", g.heroes[i].getTTL())
      log.Infof("[Hand of God] %s", message)

      g.sendEvent(message, g.heroes[i])

    }
  }
}
