package main

import "log"

func main() {
	list := []int{1, 2, 3, 4, 5}
	list2 := []int{1, 2, 3, 4, 5}
	log.Println(ProductAllButIndexWithDivision(list))
	log.Println(ProductAllButIndexWithoutDivision(list2))
}

func ProductAllButIndexWithDivision(arr []int) []int {
	product := 1
	for _, e := range arr {
		product *= e
	}
	for i := 0; i < len(arr); i++ {
		arr[i] = product / arr[i]
	}
	return arr
}

func ProductAllButIndexWithoutDivision(arr []int) []int {
	var result []int
	for i := 0; i < len(arr); i++ {
		res := 1
		for j := 0; j < len(arr); j++ {
			if j != i {
				res *= arr[j]
			}
		}
		result = append(result, res)
	}
	return result
}
