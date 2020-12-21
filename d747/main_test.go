package main

import "testing"

func TestShiftableValid(t *testing.T) {
	a := "abcde"
	b := "bcdea"
	val := IsShiftable(a, b)
	if val != true {
		t.Fail()
	}
}

func TestShiftableInvalid(t *testing.T) {
	a := "abc"
	b := "acb"
	val := IsShiftable(a, b)
	if val != false {
		t.Fail()
	}
}