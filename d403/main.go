package main

import (
	"log"
	"math/rand"
)

func rand5() int {
	return rand.Intn(5)
}

func rand7() int {
	tot := 0
	for i := 0; i < 7; i++ {
		tot += rand5()
	}
	return ((tot) % 7) + 1
}

func main() {
	var results = make(map[int]int)
	for i := 0; i < 10000; i++ {
		r := rand7()
		results[r]++
	}
	for k, v := range results {
		log.Printf("%d : %d\n", k, v)
	}
}
