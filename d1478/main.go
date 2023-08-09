package main

import "log"

type Node struct {
	Val   string
	Left  *Node
	Right *Node
}

func main() {
	tree := &Node{
		Val: "A",
		Left: &Node{
			Val: "B",
			Left: &Node{
				Val: "D",
			},
		},
		Right: &Node{
			Val: "C",
		},
	}
	log.Println(FindDeepestNode(tree).Val)
}

func FindDeepestNode(tree *Node) *Node {
	if tree == nil {
		return nil
	}
	deepestNode := tree
	queue := []*Node{tree}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node.Left != nil {
			deepestNode = node.Left
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			deepestNode = node.Right
			queue = append(queue, node.Right)
		}
	}
	return deepestNode
}
