package main

import "log"

var inverse = map[string]string{
	"0": "0",
	"1": "1",
	"8": "8",
	"6": "9",
	"9": "6",
}

func RecurseStrobogrammatic(n, length int) []string {
	if n == 0 {
		return []string{""}
	}
	if n == 1 {
		return []string{"0", "1", "8"}
	}
	middles := RecurseStrobogrammatic(n-2, length)
	result := []string{}
	for _, middle := range middles {
		if n != length {
			result = append(result, "0"+middle+"0")
		}
		result = append(result, "1"+middle+"1")
		result = append(result, "8"+middle+"8")
		result = append(result, "6"+middle+"9")
		result = append(result, "9"+middle+"6")
	}
	return result
}

func GetStrobogrammatic(n int) []string {
	return RecurseStrobogrammatic(n, n)
}

func main() {
	log.Println(GetStrobogrammatic(5))
}
