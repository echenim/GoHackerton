package main

import "fmt"

//SingleList defination for single linked list
type SingleList struct {
	head *Node
}

//Append building a linked list from an array of data
func Append(arr []string) *Node {

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

//AppendFromArray : receiver function that builds linkedlist from an array of data
func (sl *SingleList) AppendFromArray(arr []string) {
	var i int
	for i = 0; i < len(arr); i++ {
		if sl.head == nil {
			sl.head = &Node{data: arr[i]}

		}
		current := sl.head
		for current.next != nil {
			current = current.next
		}
		current.next = &Node{data: arr[i]}
	}
}

//Append : builds a linked list by appending pionter nodes
// to a receiver function
func (sl *SingleList) Append(node *Node) {
	if sl.head == nil {
		sl.head = node
		return
	}
	current := sl.head

	for current.next != nil {
		current = current.next
	}

	current.next = node
}

//Display fucntion print out the list
func Display(n *Node) {
	if n != nil {
		fmt.Printf("%v --> ", n.data)
		Display(n.next)
	}

}

//DisplayNodeWithNextMemoryAddress fucntion print out the list
func DisplayNodeWithNextMemoryAddress(n *Node, i int) {
	if n != nil {

		fmt.Printf("%v node =%+v\n", i, n)
		i++
		DisplayNodeWithNextMemoryAddress(n.next, i)
	}

}

//Display fucntion print out the list
func (sl *SingleList) Display() {
	for n := sl.head; n != nil; n = n.next {
		fmt.Printf("%v --> ", n.data)
	}
}

//DisplayNodeWithNextMemoryAddress fucntion print out the list
func (sl *SingleList) DisplayNodeWithNextMemoryAddress() {
	for n := sl.head; n != nil; n = n.next {
		fmt.Printf("e=%+v\n", n)
	}
}

//DeleteNodeFromLinkedlist from the linkedlist
func (sl *SingleList) DeleteNodeFromLinkedlist(data string) {
	if sl.head == nil {
		return
	}

	// If the head is matched then we need a new head
	if sl.head.data == data {
		// Delete head by overwriting it with next node
		sl.head = sl.head.next
		return
	}

	currentNode := sl.head
	for currentNode.next != nil {
		if currentNode.next.data == data {
			currentNode.next = currentNode.next.next
			return
		}

		currentNode = currentNode.next
	}
}
