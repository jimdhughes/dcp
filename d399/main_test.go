package main

import "testing"

func testValidInput(t *testing.T) {
	partitions := findThreePartitions([]int{3, 5, 8, 0, 8})
	if len(partitions) != 3 {
		t.Fail()
	}
}

func testNoPartitions(t *testing.T) {
	partitions := findThreePartitions([]int{1, 2, 3, 4, 5})
	if len(partitions) != 0 {
		t.Fail()
	}
}

func testShortInput(t *testing.T) {
	partitions := findThreePartitions(([]int{1}))
	if len(partitions) == 3 {
		t.Fail()
	}
}
