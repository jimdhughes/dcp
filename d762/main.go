package main

import (
	"log"
	"strings"
)

func main() {
	log.Println("hello, world")
}

func calculateMinimumDistance(firstWord, secondWord, input string) int {
	words := strings.Split(input, " ")
	firstWordIdx := -1
	secondWordIdx := -1

	for i := 0; i < len(words); i++ {
		if words[i] == firstWord {
			firstWordIdx = i
		}
		if words[i] == secondWord && firstWordIdx > -1 {
			secondWordIdx = i
		}
		return secondWordIdx - firstWordIdx
	}
	return -1
}
