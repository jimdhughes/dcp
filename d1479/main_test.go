package main

import (
	"testing"
)

func TestPartitionArrayByPivot(t *testing.T) {
	arr := []int{9, 12, 3, 5, 14, 10, 10}
	pivot := 10
	expected := []int{9, 3, 5, 10, 10, 12, 14}
	res, err := PartitionArrayByPivot(arr, pivot)
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	for i := 0; i < len(res); i++ {
		if res[i] != expected[i] {
			t.Errorf("elements do not align")
			t.FailNow()
		}
	}

}

func TestPartionArrayByPivotInPlace(t *testing.T) {
	arr := []int{9, 12, 3, 5, 14, 10, 10}
	pivot := 10
	// Note expected is not same as example but satisifes the ask
	expected := []int{5, 3, 9, 10, 10, 14, 12}
	res, err := PartionArrayByPivotInPlace(arr, pivot)
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	for i := 0; i < len(res); i++ {
		if res[i] != expected[i] {
			t.Errorf("elements do not align")
			t.FailNow()
		}
	}
}

func TestPartionArrayByPivotInPlaceLargerArray(t *testing.T) {
	arr := []int{10, 1, 5, 4, 3, 4, 5, 10, 11}
	pivot := 10
	expected := []int{5, 4, 3, 4, 5, 1, 10, 10, 11}
	res, err := PartionArrayByPivotInPlace(arr, pivot)
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	for i := 0; i < len(res); i++ {
		if res[i] != expected[i] {
			t.Errorf("elements do not align")
			t.FailNow()
		}
	}
}

func TestPartionArrayByPivotInDescendingListrArray(t *testing.T) {
	arr := []int{15, 14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	pivot := 7
	expected := []int{1, 2, 3, 4, 5, 6, 7, 15, 14, 13, 12, 11, 10, 9, 8}
	res, err := PartionArrayByPivotInPlace(arr, pivot)
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	for i := 0; i < len(res); i++ {
		if res[i] != expected[i] {
			t.Errorf("elements do not align")
			t.FailNow()
		}
	}
}
