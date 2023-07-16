package main

import "testing"

func TestFirstUseCase(t *testing.T) {
	input := "acbbac"
	expectedResult := "b"
	output := getFirstRecurringCharacterInString(input)
	if expectedResult != output {
		t.Fail()
	}
}

func TestSecondUseCase(t *testing.T) {
	input := "abcdef"
	expectedResult := ""
	output := getFirstRecurringCharacterInString(input)
	if expectedResult != output {
		t.Fail()
	}
}
