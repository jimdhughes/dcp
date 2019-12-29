package main

import "log"

var L []int = []int{1, 2, 3, 4, 5}
var cachedResults = make(map[string]int)

func sum(i, j int) int {
	key := string(i) + ":" + string(j)
	if cachedResults[key] != 0 {
		return cachedResults[key]
	}
	subL := L[i:j]
	res := 0
	for _, v := range subL {
		res += v
	}
	cachedResults[key] = res
	return res
}

func main() {
	log.Println(sum(1, 3))
	log.Println(sum(1, 4))
	log.Println(sum(1, 5))
	log.Println(sum(3, 3))
	log.Println(sum(1, 3))

}
