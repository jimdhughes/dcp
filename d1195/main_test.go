package main

import "testing"

func TestDefaultCase(t *testing.T) {
	input := "jujitsu"
	result := getShortestSubstringWithAllInputCharacters(input)
	expectedResult := "jitsu"
	if result != expectedResult {
		t.Errorf("expected \"%s\". got \"%s\"\n", expectedResult, result)
	}
}

func TestLargeCase(t *testing.T) {
	input := "aaba"
	result := getShortestSubstringWithAllInputCharacters(input)
	expectedResult := "ba"
	if result != expectedResult {
		t.Errorf("expected \"%s\". got \"%s\"\n", expectedResult, result)
	}
}
func TestStringWithAllUniqueCharacters(t *testing.T) {
	input := "abcdefghijklmnop"
	result := getShortestSubstringWithAllInputCharacters(input)
	expectedResult := input
	if result != expectedResult {
		t.Errorf("expected \"%s\". got \"%s\"\n", expectedResult, result)
	}
}

func TestSubstringContainsAllUniqueLettersOneCharacter(t *testing.T) {
	uniqueLetters := []string{"a"}
	substring := "a"
	expectedResult := true
	result := doesSubstringContainAllCharacters(uniqueLetters, substring)
	if result != expectedResult {
		t.Fail()
	}
}

func TestSubstringContainsAllUniqueLettersShouldBeFalse(t *testing.T) {
	uniqueLetters := []string{"a", "b", "c"}
	substring := "aaab"
	expectedResult := false
	result := doesSubstringContainAllCharacters(uniqueLetters, substring)
	if result != expectedResult {
		t.Fail()
	}
}
