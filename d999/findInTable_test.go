package main

import "testing"

func TestRunAllTests(t *testing.T) {
	tables := []struct {
		x int
		y int
		n int
	}{
		{6, 12, 4},
		{6, 6, 4},
		{2, 6, 0},
		{6, 36, 1},
		{10, 1, 1},
		{6, 16, 1},
		{10, 16, 3},
	}
	for _, table := range tables {
		res := CountOccurancesInMatrix(table.x, table.y)
		if res != table.n {
			t.Errorf("Wrong result for Matrix. Size: %d, value: %d. Expected: %d, got: %d\n", table.x, table.y, table.n, res)
		}
	}
}

func TestFindInMatrixFromQuestion(t *testing.T) {
	res := CountOccurancesInMatrix(6, 12)
	if res != 4 {
		t.Errorf("Expected %d got %d\n", 4, res)
	}
}

func TestFindInMatrixInTopRightCorner(t *testing.T) {
	res := CountOccurancesInMatrix(6, 6)
	if res != 4 {
		t.Fail()
	}
}

func TestNotFound(t *testing.T) {
	res := CountOccurancesInMatrix(2, 6)
	if res != 0 {
		t.Fail()
	}
}

func TestFindInMatrixBottomRightCorner(t *testing.T) {
	res := CountOccurancesInMatrix(6, 36)
	if res != 1 {
		t.Fail()
	}
}

func TestFindInMatrixOne(t *testing.T) {
	res := CountOccurancesInMatrix(10, 1)
	if res != 1 {
		t.Errorf("Expected: %d. Got: %d", 1, res)
	}
}
