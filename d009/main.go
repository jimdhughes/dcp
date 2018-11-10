package main

import "log"

func main() {
	arr := []int{5, 1, 1, 5}
	log.Printf("Max sum is: %d", CalculateSumNonAdjacent(arr, len(arr)))
}

func CalculateSumNonAdjacent(arr []int, n int) int {
	incl := arr[0]
	excl := 0
	tmp := 0
	for i := 1; i < n; i++ {
		if incl > excl {
			tmp = incl
		} else {
			tmp = excl
		}
		incl = excl + arr[i]
		excl = tmp
	}
	if incl > excl {
		return incl
	}
	return excl
}
