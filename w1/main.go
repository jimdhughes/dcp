package main

import "log"

func main() {
	var list []int
	list = append(list, 10)
	list = append(list, 15)
	list = append(list, 3)
	list = append(list, 6)

	result := 13
	hasSum := FindAnyTwoNumbersToEqualResult(result, list)
	log.Printf("List has two numbers equalling result: %t\n", hasSum)
}

func FindAnyTwoNumbersToEqualResult(result int, list []int) bool {
	var seen = make(map[int]bool)
	for _, element := range list {
		if _, ok := seen[element]; ok {
			return true
		}
		seen[result-element] = false
	}
	return false
}
