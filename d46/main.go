package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println(FindPalindrome("baabdaad"))
}

func FindPalindrome(s string) string {
	return PalindromeRecursive(s, len(s))
}

func IsPalindrome(s string) bool {
	var i, j int
	j = len(s) - 1
	log.Printf("Checking: %s\n", s)
	for i = 0; i < j; i++ {
		if s[i] != s[j] {
			return false
		}
		j--
	}
	return true
}

func PalindromeRecursive(s string, n int) string {
	if n == 1 || IsPalindrome(s) {
		return s
	}
	pLeft := PalindromeRecursive(s[1:], len(s)-1)
	pRight := PalindromeRecursive(s[:len(s)-1], len(s)-1)
	if len(pLeft) > len(pRight) {
		return pLeft
	}
	return pRight
}
