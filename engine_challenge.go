package main

import (
  "fmt"
  "math"
  "math/rand"
  "time"

  log "github.com/Sirupsen/logrus"
)

// battle function implements the battle logic
func (g *Game) CheckChallenge() {
  var h1Score, h2Score, gain int
  var message string

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

    if h1.lastBattleAt.Add(challengeCooldown).After(time.Now()) {
      // Hero fought very recently
      log.Debugf("[Challenge] %s fought very recently. Skipping", h1.HeroName)
      continue
    }

    for _, h2 := range g.heroes {
      if h1 == h2 {
        continue
      }

      if !h2.Enabled {
        continue
      }

      if h2.lastBattleAt.Add(challengeCooldown).After(time.Now()) {
        // Hero fought very recently
        log.Debugf("[Challenge] %s fought very recently. Skipping", h2.HeroName)
        continue
      }

      if heroesDistance(h1, h2) > challengeDistance {
        // Too far away
        log.Debugf("[Challenge] %s and %s are too far away.", h1.HeroName, h2.HeroName)
        continue
      }

      h1Score = rand.Intn(h1.getTotalItems()) + h1.Level
      h2Score = rand.Intn(h2.getTotalItems()) + h2.Level

      if h1Score == h2Score {
        message = fmt.Sprintf("%s and %s fought and tied.", h1.HeroName, h2.HeroName)
      } else {
        gain = int(math.Min(float64(h2.Level* challengeGainMultiplier), challengeMinGain))
        if h1Score > h2Score {
          message = fmt.Sprintf("%s, [Level:%d / Equipment:%d] has challenged %s [Level:%d / Equipment:%d] in combat and won! %d seconds are removed from %s's clock.", h1.HeroName, h1.Level, h1.getTotalItems(), h2.HeroName, h2.Level, h2.getTotalItems(), gain, h1.HeroName)
          h1.updateTTL(0 - gain)
        } else {
          message = fmt.Sprintf("%s [Level:%d / Equipment:%d] has challenged %s [Level:%d / Equipment:%d] in combat and lost! %d seconds are added to %s's clock.", h1.HeroName, h1.Level, h1.getTotalItems(), h2.HeroName, h2.Level, h2.getTotalItems(), gain, h1.HeroName)
          h1.updateTTL(gain)
        }
      }

      h1.lastBattleAt = time.Now()
      h2.lastBattleAt = time.Now()

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
