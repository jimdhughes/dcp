package main

import "fmt"

func main() {
	arr := []int{9, 11, 8, 5, 7, 10}
	fmt.Printf("%d\n", MaxProfit(arr))
}

// MaxProfit calculates max profit from trading
// arr - chronological input of stock prices
// return max profit per share from a single trade
func MaxProfit(arr []int) int {
	tArray := []int{}
	i, j, max := 0, 0, 0
	for i = 0; i < len(arr); i++ {
		tArray = append(tArray, 0)
		for j = i; j < len(arr); j++ {
			if arr[j]-arr[i] > tArray[i] {
				tArray[i] = arr[j] - arr[i]
			}
		}
	}
	for i = 0; i < len(tArray); i++ {
		if tArray[i] > max {
			max = tArray[i]
		}
	}
	return max
}
