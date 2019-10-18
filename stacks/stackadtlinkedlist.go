package main

import (
	"fmt"
)

//SingleList defination for single linked list
type SingleList struct {
	top  int
	head *Node
}




//Push : builds a linked list by appending pionter nodes
// to a receiver function
func (sl * SingleList) Push(node *Node) {
	if sl.head == nil {
		sl.top++
		sl.head = node
		return
	}
	current := sl.head

	for current.next != nil {
		sl.top++
		current = current.next
	}

	current.next = node
}



//Display fucntion print out the list
func (sl *SingleList) Display() {
	for n := sl.head; n != nil; n = n.next {
		fmt.Printf("%v <-- ", n.data)
	}
	fmt.Print("nil")
}

//Pop function to delete from stack
func (sl *SingleList) Pop() {
	if sl.head == nil {
		fmt.Print("stack underflow")
		return
	}

	//check if we have gotten to the last node
	if sl.head.next == nil {
		sl.head = sl.head.next
	}

	currentNode := sl.head
	for currentNode.next != nil {
		if currentNode.next.next == nil {
			currentNode.next = currentNode.next.next
			break
		}

		currentNode = currentNode.next
	}

}
