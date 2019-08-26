package main

import (
	"fmt"
)

//Node structure for linkedlist
type Node struct {
	data string
	next *Node
}

//Add function accept node pointer and data value
//as paramenters
func Add(node *Node, _data string) *Node {

	return &Node{data: _data, next: node}
}

//Print is a receiver function
//for printing all the nodes in the linkedlist
func (n Node) Print() {
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
