package main

import (
	"log"
	"sort"
)

func findMissingPositiveElement(arr []int) int {
	// sort array
	sort.SliceStable(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
	log.Println(arr)
	lastNum := -1
	count := 0
	for _, x := range arr {
		// handle negatives
		if x <= 0 {
			continue
		}
		// handle duplicates
		if lastNum == x {
			continue
		}
		if lastNum == -1 {
			lastNum = x
			continue
		}
		if lastNum != count+1 {
			return count + 1
		}
		lastNum = x
		count++
	}
	return lastNum + 1
}

func main() {
	arr := []int{3, 4, -1, 1}
	log.Println(findMissingPositiveElement(arr))
}
