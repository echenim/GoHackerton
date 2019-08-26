package main

import "fmt"

type node struct {
	data string
	next *node
}

func add(_data string) node {
	return node{data: _data}
}

func (n node) print() {

	current := n.next
	if current == nil {
		fmt.Printf("%v --> null", n.data)
	}
	if current != nil {
		for current.next != nil {
			fmt.Printf("%v -->", n.data)
			current = n.next
		}
	}

}
