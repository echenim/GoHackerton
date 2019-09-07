package main

//Node defination for single linked list
type Node struct {
	data string
	next *Node
}

//Dnode defination for double linked list
type Dnode struct {
	prev *Dnode
	data string
	next *Dnode
}
