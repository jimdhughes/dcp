package main

import "log"

func main() {
	log.Println(getFirstRecurringCharacterInString("acbbac"))
	log.Println(getFirstRecurringCharacterInString("abcdef"))
}

func getFirstRecurringCharacterInString(input string) string {
	charmap := make(map[rune]int)
	for _, v := range input {
		charmap[v]++
		if charmap[v] == 2 {
			return string(v)
		}
	}
	return ""
}
