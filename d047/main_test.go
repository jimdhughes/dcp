package main

import "testing"

func TestBuyLowSellHigh(t *testing.T) {
	arr := []int{9, 11, 8, 5, 7, 10}
	res := MaxProfit(arr)
	if res != 5 {
		t.Fail()
	}
}

func TestCantSellForHigher(t *testing.T) {
	arr := []int{5, 4, 3, 2, 1}
	res := MaxProfit(arr)
	if res != 0 {
		t.Fail()
	}
}
