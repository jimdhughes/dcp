package main

import "log"

func main() {
	arr := []int{10, 5, 7}
	arr2 := []int{10, 5, 1}
	log.Println(CanMakeNonDecreasing(arr))
	log.Println(CanMakeNonDecreasing(arr2))
}

func CanMakeNonDecreasing(arr []int) bool {
	countNonDecreasing := 0
	for i := 1; i < len(arr); i++ {
		if arr[i] < arr[i-1] {
			countNonDecreasing++
		}
	}
	return countNonDecreasing == 1
}
