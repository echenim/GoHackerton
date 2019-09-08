package main

import "fmt"

//DoubleLinkedlist linked list
type DoubleLinkedlist struct {
	head *Dnode
	tail *Dnode
}

//Append function implemetation for double linked list
func (dl *DoubleLinkedlist) Append(data string) {

	node := &Dnode{data: data}
	if dl.head == nil {
		dl.head = node
	} else {
		current := dl.tail
		current.next = node
		node.prev = dl.tail
	}

	dl.tail = node

}

//Insert adds an item _d at position _p
func (dl *DoubleLinkedlist) Insert(position int, data string) {
	//var i int
	if dl.head == nil {
		return
	}
	current := dl.head
	//If the head is matched then we need a new head
	if position == 1 {
		node := &Dnode{prev: nil, data: data, next: current.prev}
		dl.head = node

	}

	// //by setting i to start from 0 the new data will be inserted after the specified index
	// //by setting i to start from 0 and i is less than or equal to position the data will be inserted 2 place after the specified index
	// //by setting i to start from 1 and the if position-1 == i the new data will be inserted  the specified index
	// //by setting i to start from 1 and i is less than or equal to position the data will be inserted 1 place after the specified index

	// for i = 1; i < position; i++ {

	// 	if position-1 == i {
	// 		node := &Dnode{data: data, next: current.next}
	// 		current.next = node
	// 	}
	// 	current = current.next
	// }

}

func (dl *DoubleLinkedlist) display() {
	current := dl.head
	for current != nil {
		fmt.Printf("first node in the list is : %v \n", current.data)
		current = current.next
	}
	return
}
