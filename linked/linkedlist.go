package main

import (
	"fmt"
)

//Node structure for linkedlist
type Node struct {
	data int
	next *Node
}

//Add function to linkedlist
func Add(arr []int) *Node {

	var first *Node = new(Node) //there will be a runtime error if the point *Node is not initialize as new(Node)
	var last *Node = new(Node)  //there will be a runtime error if the point *Node is not initialize as new(Node)

	var i int

	for i = 0; i < len(arr); i++ {
		if i == 0 {
			values := arr[i]
			first.data = values
			first.next = nil
			last = first
		}
		if i > 0 {
			var t *Node = new(Node) //there will be a runtime error if the point *Node is not initialize as new(Node)
			t.data = arr[i]
			t.next = nil
			last.next = t
			last = t
		}

	}
	return first
}

//Print is a receiver function
//for printing all the nodes in the linkedlist
func (_n *Node) Print() {

	for current := _n; current != nil; current = current.next {
		if current.next != nil {
			fmt.Printf("  %v -->", current.data)
		}

		if current.next == nil {
			fmt.Printf("  %v ", current.data)
		}
	}

	fmt.Print(" --> null")
}

//Insert adds an item _d at position _p
func (_n Node) Insert(_p int, _d string) Node {

	return _n
}

//RemoveAt adds an item _d at position _p
func (_n Node) RemoveAt(_p int) {

}

//IndexOf returns the position of the Item _d
func (_n Node) IndexOf(_d int) {

}

//IsEmpty returns true if the list is empty
func (_n Node) IsEmpty() bool {
	state := false
	if _n.data == 0 && _n.next == nil {
		state = true
	}

	return state
}

//Size returns the linked list size
func (_n Node) Size() {

}

//Head returns the first node, so we can iterate on it
func (_n Node) Head() {

}
