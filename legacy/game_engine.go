package main

import (
	"log"
	"math"
	"sync"
	"time"
)

const (
	maxEnergyGain = 5
	maxEnergyLoss = 5
	maxLifeLoss   = 5
)

func (game *GameInfo) runEngine(wg *sync.WaitGroup) {
	wg.Add(len(game.Heros))
	for i := range game.Heros {
		go game.handleHero(&game.Heros[i], wg)
	}
}

func (game *GameInfo) handleHero(hero *Hero, wg *sync.WaitGroup) {
	var temperatureRatio float64
	var radiationRatio float64
	var energyGain float64
	var energyLoss float64
	var lifeLoss float64

	for hero.Life > 0 && game.Running {
		time.Sleep(1 * time.Second)
		radiationRatio = (float64)(game.Reading.Radiation-minRadiation) / (float64)(maxRadiation-minRadiation)
		energyLoss = radiationRatio * maxEnergyLoss
		if float64(hero.Energy)-energyLoss <= 0 {
			hero.Shield = false
		}

		if hero.Shield {
			hero.Energy = int64(math.Max(float64(hero.Energy)-math.Ceil(energyLoss), 0))
			log.Printf("Team %s: Energy -%.2f\n", hero.Name, energyLoss)
			continue
		}

		radiationRatio = (float64)(game.Reading.Radiation-minRadiation) / (float64)(maxRadiation-minRadiation)
		lifeLoss = radiationRatio * maxLifeLoss
		hero.Life = int64(math.Max(float64(hero.Life)-math.Ceil(lifeLoss), 0))

		temperatureRatio = (game.Reading.Temperature - minTemperature) / (maxTemperature - minTemperature)
		energyGain = temperatureRatio * maxEnergyGain
		hero.Energy = int64(math.Min(float64(hero.Energy)+math.Ceil(energyGain), 100))

		log.Printf("Hero %s : Life -%.2f, Energy +%.2f\n", hero.Name, lifeLoss, energyGain)
	}

	log.Println("Exiting goroutine for Hero: %s", hero.Name)
	wg.Done()
}
