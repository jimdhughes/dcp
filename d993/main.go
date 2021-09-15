package main

import (
	"log"
	"math"
)

func findMajorityElement(arr []int) int {
	arrLen := float64(len(arr))
	majority := math.Floor(arrLen / 2)

	var res map[int]int = make(map[int]int)

	for _, val := range arr {
		res[val]++
		if res[val] > int(majority) {
			return val
		}
	}
	return 0
}

func main() {
	log.Println("Finding majority element in array [1,2,1,1,3,4,0]")
	arr := []int{1, 2, 1, 1, 3, 4, 0}
	log.Println(findMajorityElement(arr))
}
