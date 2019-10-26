package main

import "fmt"

//SingleList defination for single linked list
type SingleList struct {
	front *Node
	rear  *Node
	head  *Node
}

//Push : builds a linked list by appending pionter nodes
// to a receiver function
func (sl *SingleList) Push(node *Node) {

	if sl.head == nil {
		sl.front = node
		sl.rear = sl.front
		sl.head = node
		return
	}

	current := sl.head
	for current.next != nil {
		current = current.next
	}
	current.next = node
	sl.rear = node
}

//Display fucntion print out the list
func (sl *SingleList) Display() {
	for n := sl.head; n != nil; n = n.next {
		fmt.Printf("%v <-- ", n.data)
	}
	fmt.Print("nil")
}
