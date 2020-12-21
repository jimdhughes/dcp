package main

import "testing"

func TestRandomBtree(t *testing.T) {
	arbitraryDepth := 0
	for node := generate() ; node != nil; node = traverseNode(node) {
		arbitraryDepth++
	}
	if arbitraryDepth == 0 {
		t.Fail()
	}
}

func traverseNode(n *Node) *Node {
	if n.left != nil {
		return n.left
	}
	if n.right != nil {
		return n.right
	}
	return nil
}