package main

import (
	"fmt"
	"strings"
)

func LongestSubstringWithKCharacters(s string, k int) string {
	var sub string
	var tmp string
	var visited string
	start := 0
	idx := 0

	for {
		if idx == len(s) {
			break
		}
		c := s[idx : idx+1] // get the character to process
		tmp = tmp + c       // append c to tmp
		if len(visited) < k {
			if !strings.Contains(visited, c) {
				visited = visited + c
			}
		} else {
			if !strings.Contains(visited, c) {
				// reset condition
				tmp = ""
				visited = ""
				start = start + 1
				idx = start - 1 // init idx to be start -1 since we always increment idx at the end
			}
		}
		if len(tmp) > len(sub) {
			sub = tmp
		}
		idx = idx + 1
	}
	return sub
}

func main() {
	fmt.Println(LongestSubstringWithKCharacters("abcba", 2))
	fmt.Println(LongestSubstringWithKCharacters("ABBCCCDDDD", 3))
	fmt.Println(LongestSubstringWithKCharacters("ABBCCCDDDD", 2))
	fmt.Println(LongestSubstringWithKCharacters("ABBCCCDDDD", 1))
	fmt.Println(LongestSubstringWithKCharacters("ABBCCCDDDD", 0))
}
