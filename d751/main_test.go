package main

import (
	"testing"
)

func TestInorderTraversal(t *testing.T) {
	result := make([]int,0)
	node := Node{
		Value:20,
		Left: &Node {
			Value: 8,
			Left: &Node {
				Value: 4,
			},
			Right: &Node {
				Value: 12,
				Left: &Node{
					Value: 10,
				},
				Right: &Node {
					Value: 14,
				},
			},
		},
		Right: &Node {
				Value: 22,
		},
	}
	expected:= []int{4,8,10,12,14,20,22}
	InorderTraversal(node, &result)
	if len(expected) != len(result) {
		t.Fail()
	}
	for idx := range(expected) {
		if result[idx] != expected[idx] {
			t.Fail()
		}
	}
}