package main

import "testing"

func TestHasMajoritySingleEntry(t *testing.T) {
	arr := []int{1}
	res := findMajorityElement(arr)
	if res != 1 {
		t.Fail()
	}
}

func TestHasMajorityThreeEntries(t *testing.T) {
	arr := []int{0, 1, 1}
	res := findMajorityElement(arr)
	if res != 1 {
		t.Fail()
	}
}

func TestHasMultipleMajorityReturnFirst(t *testing.T) {
	arr := []int{0, 0, 1, 1}
	res := findMajorityElement(arr)
	if res != 0 {
		t.Fail()
	}
}

func TestHasMultipleMajorityLargestLast(t *testing.T) {
	arr := []int{0, 1, 2, 3, 4, 4, 4, 4, 4}
	res := findMajorityElement(arr)
	if res != 4 {
		t.Fail()
	}
}
