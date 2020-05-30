package main

import "testing"

func TestFindElementInArray(t *testing.T) {
	arr := []int{13, 18, 25, 2, 8, 10}
	idx := FindElementInArray(arr, 8, 0, len(arr)-1)
	if idx != 4 {
		t.Fail()
	}
}

func TestFindElementInArray_SingleElementNoMatch(t *testing.T) {
	arr := []int{1}
	idx := FindElementInArray(arr, 2, 0, len(arr)-1)
	if idx != -1 {
		t.Fail()
	}
}
