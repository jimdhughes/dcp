package main

import "testing"

func TestScenarioOne(t *testing.T) {
	arr := []int{3, 4, -1, 1}
	val := findMissingPositiveElement(arr)
	if val != 2 {
		t.Fail()
	}
}

func TestScenarioTwo(t *testing.T) {
	arr := []int{2, 1, 0}
	val := findMissingPositiveElement(arr)
	if val != 3 {
		t.Fail()
	}
}
