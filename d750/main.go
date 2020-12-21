package main

import (
	"math/rand"
)

// Node is a simple tree node
type Node struct {
	val int
	left *Node
	right *Node
}

func generate() *Node {
	node := &Node{
		val: 0,
	}
	// randomly generate left node
	rRand := rand.Intn(100)
	if rRand > 50 {
		node.right = generate()
	}
	// randomly generate right node
	lRand := rand.Intn(100)
	if lRand > 50 {
		node.left = generate()
	}
	return node
}