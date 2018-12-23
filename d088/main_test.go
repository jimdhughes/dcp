package main

import "testing"

func Test10DividiedBy5(t *testing.T) {
	res := DivideIntegers(10, 5)
	if res != 2 {
		t.Fail()
	}
}

func TestDivisorLargerThanDividend(t *testing.T) {
	res := DivideIntegers(5, 10)
	if res != 0 {
		t.Fail()
	}
}

func TestUnevenDivision(t *testing.T) {
	res := DivideIntegers(8, 7)
	if res != 1 {
		t.Fail()
	}
}

func TestDivideBy0(t *testing.T) {
	res := DivideIntegers(10, 0)
	if res != 0 {
		t.Fail()
	}
}
