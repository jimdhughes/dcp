package main

import "testing"

func TestCanMakeNonDecreasing_True(t *testing.T) {
	arr := []int{10, 5, 7}
	res := CanMakeNonDecreasing(arr)
	if res != true {
		t.Fail()
	}
}

func TestCanMakeNonDecreasing_False(t *testing.T) {
	arr := []int{10, 5, 1}
	res := CanMakeNonDecreasing(arr)
	if res != false {
		t.Fail()
	}
}

func TestCanMakeNonDecreasingSingleElementArray(t *testing.T) {
	arr := []int{1}
	res := CanMakeNonDecreasing(arr)
	if res != false {
		t.Fail()
	}
}

func TestCanMakeNonDecreasingLongArray_true(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 11, 10}
	if CanMakeNonDecreasing(arr) != true {
		t.Fail()
	}
}

func TestCanMakeNonDecreasingLongArray_False(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 12, 10, 13, 11}
	if CanMakeNonDecreasing(arr) != false {
		t.Fail()
	}
}
