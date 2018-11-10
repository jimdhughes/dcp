package main

import "log"

func main() {
	message := "1111"
	runes := []rune(message)
	log.Printf("Valid Combinations: %d", CalculateValidMessageCombinations(runes, len(message)))
}

func CalculateValidMessageCombinations(runes []rune, n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	count := 0
	// handle single digit case
	if runes[n-1] > '0' {
		count += CalculateValidMessageCombinations(runes[:n-1], n-1)
	}
	// handle double digit case
	if runes[n-2] == '1' || runes[n-2] == '2' && runes[n-1] < '7' {
		count += CalculateValidMessageCombinations(runes[:n-2], n-2)
	}
	return count

}
