package main

import "testing"

func TestFindInMatrixFromQuestion(t *testing.T) {
	res := FindInMatrix(6, 12)
	if res != 4 {
		t.Errorf("Expected %d got %d\n", 4, res)
	}
}

func TestFindInMatrixInTopRightCorner(t *testing.T) {
	res := FindInMatrix(6, 6)
	if res != 3 {
		t.Fail()
	}
}

func TestNotFound(t *testing.T) {
	res := FindInMatrix(2, 6)
	if res != 0 {
		t.Fail()
	}
}

func TestFindInMatrixBottomRightCorner(t *testing.T) {
	res := FindInMatrix(6, 36)
	if res != 1 {
		t.Fail()
	}
}

func TestFindInMatrixOne(t *testing.T) {
	res := FindInMatrix(10, 1)
	if res != 1 {
		t.Errorf("Expected: %d. Got: %d", 1, res)
	}
}
