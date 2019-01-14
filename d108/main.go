package main

import (
	"log"
	"strings"
)

func main() {
	log.Println(ShiftTester("BCDEFGHIJKLMNOPA", "ABCDEFGHIJKLMNOP"))
}

func ShiftTester(a, b string) bool {
	if len(a) != len(b) {
		return false
	}
	var temp string
	for i := 1; i < len(a); i++ {
		temp = a[i:] + a[0:i]
		if strings.Compare(temp, b) == 0 {
			return true
		}
	}
	return false

}
