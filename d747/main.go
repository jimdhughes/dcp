package main

import "strings"

// IsShiftable will determine if string b can be shifted
// by some number to match string a
func IsShiftable(a, b string) bool {
	if len(a) != len(b) {
		return false
	}
	if a == b {
		return true
	}
	b = b + b
	idx := strings.Index(b, a)
	if idx >= 0  {
		return true
	}
	return false
}