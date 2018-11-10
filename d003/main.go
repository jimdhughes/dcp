package main

import (
	"log"
	"strings"
)

type Node struct {
	value       string
	left, right *Node
}

//Serialize Nodes into a string
func (n *Node) serialize() string {
	paths := []string{}
	SerializeRecursive(n, &paths)
	return strings.Join(paths, ",")
}

func SerializeRecursive(n *Node, paths *[]string) {
	if n == nil {
		*paths = append(*paths, "")
		return
	}
	*paths = append(*paths, n.value)
	SerializeRecursive(n.left, paths)
	SerializeRecursive(n.right, paths)
}

func (n *Node) deserialize(serial string) {
	arr := strings.Split(serial, ",")
	node := &Node{}
	DeserializeRecursive(node, arr)
}

func DeserializeRecursive(node *Node, arr []string) {
	if len(arr) == 0 {
		return
	}
	el := arr[0]
	arr = arr[1:]
	if el == "" {
		return
	}
	node.left = &Node{value: el}
	DeserializeRecursive(node, arr)
	node.right = &Node{value: el}
	DeserializeRecursive(node, arr)
}

func main() {
	node := Node{
		value: "root",
		left: &Node{
			value: "left",
			left: &Node{
				value: "left.left",
			},
		},
		right: &Node{
			value: "right",
		},
	}
	serial := node.serialize()
	log.Println(serial)
	node.deserialize(serial)
	if node.left.left.value == "left.left" {
		log.Println("Successful")
		log.Println(node.serialize())
		return
	}
	log.Println("Fail")

}
