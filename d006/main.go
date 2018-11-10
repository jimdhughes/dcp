package main

import (
	"log"
	"unsafe"
)

type Node struct {
	value int
	npx   *Node
}

func XOR(prev, next *Node) *Node {
	return (*Node)(unsafe.Pointer(uintptr(unsafe.Pointer(prev)) ^ uintptr(unsafe.Pointer(next))))
}

func add(head **Node, value int) {
	newNode := &Node{value: value}
	newNode.npx = XOR(*head, nil)
	if *head != nil {
		next := XOR((*head).npx, nil)
		(*head).npx = XOR(newNode, next)
	}
	*head = newNode
}

func get(head *Node, index int) *Node {
	curr := head
	var prev, next *Node
	for i := 0; i < index; i++ {
		next = XOR(prev, curr.npx)
		prev = curr
		curr = next
	}
	return curr
}

func main() {
	var node *Node
	add(&node, 10)
	add(&node, 12)
	i := get(node, 1)
	if i != nil {
		log.Printf("Value is: %d", i.value)
		return
	}
	log.Printf("Node is nil")
}
