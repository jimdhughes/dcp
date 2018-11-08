package main

import "testing"

var testString = "ABBCCCDDDD"

func TestSubstringAllSingleCharacters(t *testing.T) {
	res := LongestSubstringWithKCharacters("aaaaa", 3)
	if res != "aaaaa" {
		t.Fail()
	}
}

func TestSubstringShouldReturnDDD(t *testing.T) {
	res := LongestSubstringWithKCharacters(testString, 1)
	if res != "DDDD" {
		t.Fail()
	}
}
func TestSubstringShouldReturnCCCDDDD(t *testing.T) {
	res := LongestSubstringWithKCharacters(testString, 2)
	if res != "CCCDDDD" {
		t.Fail()
	}
}
func TestSubstringShouldReturnBBCCCDDDD(t *testing.T) {
	res := LongestSubstringWithKCharacters(testString, 3)
	if res != "BBCCCDDDD" {
		t.Fail()
	}
}

func TestSubstringFromEmptyString(t *testing.T) {
	res := LongestSubstringWithKCharacters("", 3)
	if res != "" {
		t.Fail()
	}
}
