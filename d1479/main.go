package main

import (
	"errors"
	"log"
)

func main() {
	log.Println("D1479")

}

// Naive implementation
func PartitionArrayByPivot(arr []int, pivot int) ([]int, error) {
	var smallerElements []int
	var largerElements []int
	var equalElements []int
	if len(arr) == 0 {
		return []int{}, errors.New("empty array")
	}
	for _, v := range arr {
		if v < pivot {
			smallerElements = append(smallerElements, v)
		} else if v > pivot {
			largerElements = append(largerElements, v)
		} else {
			equalElements = append(equalElements, v)
		}
	}
	return append(append(smallerElements, equalElements...), largerElements...), nil
}

func PartionArrayByPivotInPlace(arr []int, pivot int) ([]int, error) {
	if len(arr) == 0 {
		return arr, errors.New("empty array")
	}
	var start, end int = 0, len(arr) - 1

	for start <= end {
		// if the value equal to the pivot, do nothing
		if arr[start] == pivot {
			start++
		}
		// if the value is greater than pivot, swap with the end
		if arr[start] > pivot {
			temp := arr[start]
			// remove the element from the array
			arr = append(arr[:start], arr[start+1:]...)
			arr = append(arr, temp)
			end--
		}
		// if the value is equal to pivot, swap with the start
		if arr[start] < pivot {
			temp := arr[start]
			arr = append(arr[:start], arr[start+1:]...)
			arr = append([]int{temp}, arr...)
			start++
		}
	}
	return arr, nil
}
