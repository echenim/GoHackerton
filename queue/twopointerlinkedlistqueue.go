package main

import "fmt"

//SingleList defination for single linked list
type SingleList struct {
	front *Node
	rear  *Node
	head  *Node
}

//Enqueue : builds a linked list by appending pionter nodes
// to a receiver function
func (sl *SingleList) Enqueue(node *Node) {

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

//Dequeue : builds a linked list by appending pionter nodes
// to a receiver function
func (sl *SingleList) Dequeue() {
	if sl.head == nil {
		fmt.Println("\nQueue is empty\n")
	}

	//check if we have gotten to the last node
	if sl.head.next == nil {
		sl.head = sl.head.next
		fmt.Println("\nQueue is empty T")
	}

	if sl.head != nil {
		sl.head = sl.head.next
		sl.front = &Node{data: sl.head.data, next: nil}
	}

}
