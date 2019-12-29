package main

import (
	"log"
)

type Node struct {
	Val  int
	Next *Node
}

func (l *Node) Push(i int) *Node {
	n := &Node{Val: i}
	n.Next = l
	return n
}

func (l *Node) RemoveNode(k int) (*Node, error) {
	d := &Node{}
	d.Next = l
	firstNode := d
	secondNode := d
	for i := 0; i < k+1; i++ {
		firstNode = firstNode.Next
	}
	log.Println(firstNode)
	log.Println(d)
	if firstNode.Next == nil {
		log.Println("removing first")
		d.Next = d.Next.Next
		return d.Next, nil
	}
	for firstNode.Next != nil {
		firstNode = firstNode.Next
		secondNode = secondNode.Next
	}
	secondNode.Next = secondNode.Next.Next
	return d.Next, nil
}

func (l *Node) PrintList() {
	t := l
	for t != nil {
		log.Println(t.Val)
		t = t.Next
	}
}

func main() {
	head := &Node{Val: 1}
	head = head.Push(2)
	head = head.Push(3)
	head = head.Push(4)
	head = head.Push(5)
	head = head.Push(6)
	head = head.Push(7)
	head.PrintList()
	head, _ = head.RemoveNode(0)
	head.PrintList()
}
