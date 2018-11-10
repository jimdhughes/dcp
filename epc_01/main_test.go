package main

import (
	"testing"
)

func TestSingleCharacterArray(t *testing.T) {
	s := "a"
	res := LongestNonRepeatingSubstring(s)
	if res != s {
		t.Fail()
	}
}

func TestTwoCharacterMaxLengthArray(t *testing.T) {
	s := "aaaaaaaab"
	res := LongestNonRepeatingSubstring(s)
	if res != "ab" {
		t.Fail()
	}
}

func TestTwoCharacterMaxLengthArrayStart(t *testing.T) {
	s := "abaaaaaa"
	res := LongestNonRepeatingSubstring(s)
	if res != "ab" {
		t.Fail()
	}
}

func TestNonRepeatingCharacterArray(t *testing.T) {
	s := "ABCDEFG"
	res := LongestNonRepeatingSubstring(s)
	if s != res {
		t.Fail()
	}
}

func TestEmptyArray(t *testing.T) {
	s := ""
	res := LongestNonRepeatingSubstring(s)
	if s != res {
		t.Fail()
	}
}

func TestCharacterSequenceInMiddle(t *testing.T) {
	s := "HELLOEPCORRRR"
	res := LongestNonRepeatingSubstring(s)
	if res != "LOEPC" {
		t.Fail()
	}
}
func TestLongestSequenceAtStart(t *testing.T) {
	s := "ABCDEFGHIAAAAAAAAAAA"
	res := LongestNonRepeatingSubstring(s)
	if res != "ABCDEFGHI" {
		t.Fail()
	}
}
