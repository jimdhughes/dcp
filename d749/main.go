package main

// LongestPalindrome checks a string to find it's longest palindrome
func LongestPalindrome(input string) string {
	if len(input) == 1 {
		return input
	}
	if IsPalindrome(input) {
		return input
	}
	right := LongestPalindrome(input[1:])
	left := LongestPalindrome(input[0:len(input)-1])
	if len(right) > len(left) {
		return right
	}
	return left
}

// IsPalindrome checks that a string is a palindrome
func IsPalindrome(input string) bool {
	i:=0
	j:=len(input)-1
	for i <= j {
		if input[i] != input[j] {
			return false
		}
		i++
		j--
	}
	return true
}