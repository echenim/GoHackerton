package main

//Bnode defination for each tree
type Bnode struct {
	data       byte
	leftChild  *Bnode
	rightChild *Bnode
}
