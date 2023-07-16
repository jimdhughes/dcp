package main

import (
	"log"
)

func findMajorityElement(arr []int) int {
	m := make(map[int]int)
	for _, v := range arr {
		m[v]++
	}
	for k, v := range m {
		if v > len(arr)/2 {
			return k
		}
	}
	return -1
}


func main() {
	log.Println("Finding majority element in array [1,2,1,1,3,4,0]")
	arr := []int{1, 2, 1, 1, 3, 4, 0}
	log.Println(arr)
}
