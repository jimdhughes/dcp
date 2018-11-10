package main

import (
	"fmt"
)

// LongestNonRepeatingSubstring finds the longrest non-repeating substring in a given string
func LongestNonRepeatingSubstring(s string) string {
	var i, maxLength, maxLengthStartIndex int = 0, 0, 0
	charmap := make(map[byte]int)
	for j := 0; j < len(s); j++ {
		charmap[s[j]]++
		for charmap[s[j]] == 2 && i < j {
			charmap[s[i]]--
			i++
		}
		if maxLength < (j - i + 1) {
			maxLength = (j - i + 1)
			maxLengthStartIndex = i
		}
	}
	return s[maxLengthStartIndex : maxLengthStartIndex+maxLength]
}

func main() {
	fmt.Println(LongestNonRepeatingSubstring("aabcabcd"))
}
