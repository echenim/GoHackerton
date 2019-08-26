package main

import (
	"fmt"
)

//Node structure for linkedlist
type Node struct {
	data string
	next *Node
}

func add(node *Node, _data string) *Node {

	return &Node{data: _data, next: node}
}

func (n Node) print() {
	fmt.Printf("%v ", n.data)
	current := n.next
	if current.next == nil {
		//	fmt.Printf("%v --> null", n.data)
	} else {
		for current != nil {
			fmt.Printf(" --> %v", current.data)
			current = current.next

		}
	}
	fmt.Print(" --> null")
}
