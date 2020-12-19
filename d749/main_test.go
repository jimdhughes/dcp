package main

import "testing"

func TestIsPalindromeEvenValid(t *testing.T) {
	res := IsPalindrome("abba")
	if res == false {
		t.Fail()
	}
}

func TestIsPalindromeOddValid(t *testing.T) {
	res := IsPalindrome("abcba")
	if res == false {
		t.Fail()
	}
}

func TestIsPalindromeEvenInvalid(t *testing.T) {
	res := IsPalindrome("aabb")
	if res == true {
		t.Fail()
	}
}

func TestIsPalindromeOddInvalid(t *testing.T) {
	res := IsPalindrome("aabba")
	if res == true {
		t.Fail()
	}
}

func TestLongestPalindromeFullPalindromeValid(t *testing.T) {
	input := "abcba"
	expected := input
	res := LongestPalindrome(input)
	if res != input {
		t.Errorf("Expected %s, Got %s\n", expected, res)
		t.Fail()
	}
}

func TestLongestPalindromeLeftValid(t *testing.T) {
	input := "abcbaz"
	expected := "abcba"
	res := LongestPalindrome(input)
	if res != expected {
		t.Errorf("Expected %s, Got %s\n", expected, res)
		t.Fail()
	}
}

func TestLongestPalindromeRightValid(t *testing.T) {
	input := "zabcba"
	expected := "abcba"
	res := LongestPalindrome(input)
	if res != expected {
		t.Errorf("Expected %s, Got %s\n", expected, res)
		t.Fail()
	}
}

func TestLongestPalindromeInMiddleValid(t *testing.T) {
	input := "zxyabcbazxy"
	expected := "abcba"
	res := LongestPalindrome(input)
	if res != expected {
		t.Errorf("Expected %s, Got %s\n", expected, res)
		t.Fail()
	}
}

func TestLongestPalindromeNoPalindromeValid(t *testing.T) {
	input := "abcdefghijklmnop"
	
	res := LongestPalindrome(input) 
	if len(res) > 1 {
		t.Error("Expected a 1 character result")
		t.Fail()
	}
}

func TestLongestPalindromeQuestionInputValid(t *testing.T) {
	input := "aabcdcb"
	expected := "bcdcb"
	res := LongestPalindrome(input)
	if res != expected {
	
		t.Fail()
	}
}

func TestLongestPalindromeQuestionInputTwoValid(t *testing.T) {
	input := "bananas"
	expected := "anana"
	res := LongestPalindrome(input)
	if res != expected {
			t.Errorf("Expected %s, Got %s\n", expected, res)
			t.Fail()
	}
}