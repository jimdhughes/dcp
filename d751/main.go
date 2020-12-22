package main

type Node struct {
	Value int
	Left *Node
	Right *Node
}

func InorderTraversal(node Node, result *[]int)  {
	if node.Left != nil {
		InorderTraversal(*node.Left, result)
	}
	*result = append(*result, node.Value)
	if node.Right != nil {
		InorderTraversal(*node.Right, result)
	}
}