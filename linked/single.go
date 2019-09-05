package main

import (
	"fmt"
	"strconv"
)

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
	fmt.Print("nil")
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

//Insert adds an item _d at position _p
func (sl *SingleList) Insert(position int, data string) {
	var i int
	if sl.head == nil {
		return
	}
	current := sl.head
	//If the head is matched then we need a new head
	if position == 1 {
		node := &Node{data: data, next: current}
		sl.head = node

	}

	//by setting i to start from 0 the new data will be inserted after the specified index
	//by setting i to start from 0 and i is less than or equal to position the data will be inserted 2 place after the specified index
	//by setting i to start from 1 and the if position-1 == i the new data will be inserted  the specified index
	//by setting i to start from 1 and i is less than or equal to position the data will be inserted 1 place after the specified index

	for i = 1; i < position; i++ {

		if position-1 == i {
			node := &Node{data: data, next: current.next}
			current.next = node
		}
		current = current.next
	}

}

func (sl *SingleList) size() int {
	i := 0
	if sl.head == nil {
		return i
	}
	node := sl.head
	for node != nil {
		i++
		node = node.next
	}

	return i
}

//ReversingData reversing linked list by reversing the data
func (sl *SingleList) ReversingData(size int) {
	arr := make([]string, size)
	var ai int

	for current := sl.head; current != nil; current = current.next {
		arr[ai] = current.data
		ai++
	}

	for current := sl.head; current != nil; current = current.next {
		ai--
		current.data = arr[ai]

	}

}

//ReversingLinks reversing linked list by reversing the data
func (sl *SingleList) ReversingLinks() {
	var q *Node
	var r *Node
	current := sl.head
	for current != nil {
		r = q
		q = current
		current = current.next
		q.next = r

	}

	sl.head = q
}

//Concatenation concat two linked list
func Concatenation(sl *SingleList, sll *SingleList) {
	current := sl.head
	for current.next != nil {
		current = current.next
	}
	current.next = sll.head
}

//Merge fucntion for list a and b without allowing duplicates in both list
func Merge(sl *SingleList, sll *SingleList) *SingleList {
	mergelist := &SingleList{}
	a := sl.head
	b := sll.head

	for a.next != nil && b.next != nil {
		ai, _ := strconv.ParseInt(a.data, 10, 0)
		bi, _ := strconv.ParseInt(b.data, 10, 0)
		if ai < bi {
			mergelist.Append(&Node{data: a.data, next: nil})
			a = a.next
		}

		if ai > bi {
			mergelist.Append(&Node{data: b.data, next: nil})
			b = b.next
		}

	}

	if a.next != nil && b.next == nil {
		mergelist.Append(&Node{data: a.data, next: nil})
		a = a.next
	}

	if a.next == nil && b.next != nil {
		mergelist.Append(&Node{data: b.data, next: nil})
		b = b.next
	}
	return mergelist
}

//MergeWithDuplicates fucntion for linked list sl and sll and allow duplicates in both list to
func MergeWithDuplicates(sl *SingleList, sll *SingleList) *SingleList {
	mergelist := &SingleList{}
	a := sl.head
	b := sll.head

	for a.next != nil && b.next != nil {
		ai, _ := strconv.ParseInt(a.data, 10, 0)
		bi, _ := strconv.ParseInt(b.data, 10, 0)
		if ai <= bi {
			mergelist.Append(&Node{data: a.data, next: nil})
			a = a.next
		}

		if ai > bi {
			mergelist.Append(&Node{data: b.data, next: nil})
			b = b.next
		}

	}

	if a.next != nil && b.next == nil {
		mergelist.Append(&Node{data: a.data, next: nil})
		a = a.next
	}

	if a.next == nil && b.next != nil {
		mergelist.Append(&Node{data: b.data, next: nil})
		b = b.next
	}
	return mergelist
}
