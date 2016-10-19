package main

import (
  "fmt"
  "math/rand"

  log "github.com/Sirupsen/logrus"
)

// findItem generates a new item for the hero when they level up
// and notifies the player about the item found
func (g *Game) findItem(hero *Hero) {

  items := []string{"weapon", "tunic", "shield", "leggings", "ring", "gloves", "boots", "helm", "charm", "amulet"}

  findChance := []float32{100.00, 91.93227152, 84.51542547, 77.69695042, 71.42857143, 65.66590823, 60.36816105, 55.49782173,
    51.02040816, 46.90422016, 43.12011504, 39.64130124, 36.44314869, 33.5030144, 30.80008217, 28.31521517, 26.03082049,
    23.93072457, 22.00005869, 20.22515369, 18.59344321, 17.0933747, 15.71432764, 14.44653835, 13.28103086, 12.20955335,
    11.22451974, 10.31895597, 9.486450616, 8.721109539, 8.017514101, 7.370682832, 6.776036155, 6.229363956, 5.726795786,
    5.264773452, 4.840025825, 4.449545683, 4.090568419, 3.760552466, 3.457161303, 3.178246916, 2.921834585, 2.686108904,
    2.469400931, 2.270176369, 2.087024703, 1.918649217, 1.763857808, 1.621554549, 1.490731931}

  var itemType string
  var newItemLevel int
  var itemFoundChance float32
  found := false

  for i := hero.Level; i > 0; i-- {

    if i > 50 {
      //After Hero Level of 50, has a 1% chance to find an item.
      itemFoundChance = 1.0
    } else {
      itemFoundChance = findChance[i]
    }

    //Start with highest Level Item and subtract a level as it misses the chance
    if rand.Intn(100) <= int(itemFoundChance) {
      // Item found!
      found = true
      itemGainPercentage := float64(rand.Intn(100))
      newItemLevel = int(float64(i) + (float64(i) * (itemGainPercentage / 100)))
      itemType = items[rand.Intn(10)]
      break
    }
  }

  if found {
    currentItemLevel := hero.getItemLevel(itemType)

    log.Debugf("Item Found: %s | Hero Level: %d | Current Item Level: %d | New Item Level: %d", itemType, hero.Level, currentItemLevel, newItemLevel)

    var message string
    verb := "is"

    if itemType == "leggings" || itemType == "gloves" || itemType == "boots" {
      verb = "are"
    }

    if newItemLevel > currentItemLevel {
      // Replace the current item value with the new one
      hero.updateItem(itemType, newItemLevel)
      message = fmt.Sprintf("You found a level %d %s! Your current %s %s only level %d, so it seems luck is with you!", newItemLevel, itemType, itemType, verb, currentItemLevel)
    } else {
      // Message back to player that current item level is better
      message = fmt.Sprintf("You found a level %d %s! Your current %s %s level %d, so it seems Luck is against you. You toss the %s", newItemLevel, itemType, itemType, verb, currentItemLevel, itemType)
    }

    g.sendEvent(message, hero)
  } else {
    log.Debugf("No items found for Hero: %s (ID: %d) | Level: %d", hero.HeroName, hero.id, hero.Level)
  }
}
