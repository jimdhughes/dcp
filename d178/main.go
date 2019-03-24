package main

import (
	"log"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	iterations := 1000000
	runScenario(5, 6, iterations)
	runScenario(5, 5, iterations)
}

func runScenario(value1, value2, iterations int) {
	log.Printf("Running %d iterations\n", iterations)
	cost := 0
	min := 0
	max := 0
	for i := 0; i < iterations; i++ {
		iterCost := playGame(value1, value2)
		if iterCost < min || min == 0 {
			min = iterCost
		}
		if iterCost > max || max == 0 {
			max = iterCost
		}
		cost += iterCost
	}

	log.Printf("Scenairo: %d, %d\n", value1, value2)
	log.Printf("Results:\nMin: %d\nMax: %d\nAvg: %0.2f\n", min, max, float64(cost)/float64(iterations))
}

func playGame(value1, value2 int) int {
	cost := 1
	lastRoll := 0
	for {
		roll := rollDie()
		if lastRoll == value1 && roll == value2 {
			return cost
		}
		lastRoll = roll
		cost = cost + 1
	}
}

func rollDie() int {
	return rand.Intn(6) + 1
}
