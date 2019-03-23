package main

import (
	"log"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	iterations := 1000000
	log.Printf("Running %d iterations\n", iterations)
	costGame1 := 0
	for i := 0; i < iterations; i++ {
		costGame1 = costGame1 + playGame(5, 6)
	}
	costGame2 := 0
	for i := 0; i < iterations; i++ {
		costGame2 = costGame2 + playGame(5, 5)
	}
	log.Printf("Avg Game 1 cost: %0.2f\n", float64(costGame1)/float64(iterations))
	log.Printf("Avg Game 2 cost: %0.2f\n", float64(costGame2)/float64(iterations))
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
