package main

import "log"

type Node struct {
	Val         int
	Left, Right *Node
}

func CountUnivalTrees(n *Node) int {
	if n.Left == nil && n.Right == nil {
		return 1
	} else if n.Left.Val == n.Right.Val {
		return 1 + CountUnivalTrees(n.Left) + CountUnivalTrees(n.Right)
	} else {
		return 0 + CountUnivalTrees(n.Left) + CountUnivalTrees(n.Right)
	}
}

func main() {
	n := &Node{Val: 0}
	n.Left = &Node{Val: 1}
	n.Right = &Node{Val: 0}
	n.Right.Left = &Node{Val: 1}
	n.Right.Right = &Node{Val: 0}
	n.Right.Left.Left = &Node{Val: 1}
	n.Right.Left.Right = &Node{Val: 1}
	log.Println(CountUnivalTrees(n))
}
