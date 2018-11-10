package main

import "testing"

func TestShortPalindrome(t *testing.T) {
	s := "aba"
	res := FindPalindrome("aba")
	if res != s {
		t.Errorf("Expected %s got %s", s, res)
	}
}

func TestNoPalindrome(t *testing.T) {
	s := "abcdefg"
	res := FindPalindrome(s)
	if len(res) > 1 {
		t.Errorf("Expected no palindrome. Found %s", res)
	}
}

func TestPalindromeIsFullString(t *testing.T) {
	s := "tattarrattat"
	res := FindPalindrome(s)
	if res != s {
		t.Errorf("Expected %s got %s", s, res)
	}
}

func TestLongWordNoPalindrome(t *testing.T) {
	s := "abcdefghijklmnopqrstuvwxyz"
	res := FindPalindrome(s)
	if len(res) > 1 {
		t.Errorf("Expected no palindrome. Found %s", res)
	}
}

func testEvenPalindrome(t *testing.T) {
	s := "abba"
	res := FindPalindrome(s)
	if res != s {
		t.Errorf("Expected %s got %s", "abba", res)
	}
}

func TestOddPalindrome(t *testing.T) {
	s := "aabaa"
	res := FindPalindrome(s)
	if res != s {
		t.Errorf("Expected %s got %s", "aabaa", res)
	}
}

func TestEvenPalindromeMidString(t *testing.T) {
	s := "abcdefgabbaabcdefg"
	res := FindPalindrome(s)
	if res != "abba" {
		t.Errorf("Expected %s got %s", "abba", res)
	}
}

func TestOddPalindromeMidString(t *testing.T) {
	s := "abcdefgabcbaaxcdefg"
	res := FindPalindrome(s)
	if res != "abcba" {
		t.Errorf("Expected %s got %s", "abcba", res)
	}
}
