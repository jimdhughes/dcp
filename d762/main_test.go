package main

import "testing"

func testCalculateMinDistanceExpectTwo(t *testing.T) {
	firstWord := "hello"
	secondWord := "world"
	input := "dog cat hello cat dog dog hello cat world"
	res := calculateMinimumDistance(firstWord, secondWord, input)
	if res != 1 {
		t.Fail()
	}
}

func testCalculateMinDistanceExpectNegativeOne(t *testing.T) {
	firstWord := "hello"
	secondWord := "world"
	input := "world hello"
	res := calculateMinimumDistance(firstWord, secondWord, input)
	if res != -1 {
		t.Fail()
	}
}
