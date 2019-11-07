package main

import "fmt"

//Btree defination
type Btree struct {
	root *Bnode
}

//Insert function for creating tree
func (t *Btree) Insert(node *Bnode) *Btree {

	if t.root == nil {
		t.root = node
	} else {
		t.root.insert(node)
	}
	return t
}

//insert fuction that will be called recursively to fill the tree node
func (b *Bnode) insert(node *Bnode) {
	if b == nil {
		return
	} else if node.data > b.data {

		if b.leftChild == nil {
			b.leftChild = node
		} else {
			b.leftChild.insert(node)
		}

	} else {
		if b.rightChild == nil {
			b.rightChild = node
		} else {
			b.rightChild.insert(node)
		}
	}
}

//Print fucntion for displaying the tree nodes and branch
func (t *Btree) Print(node *Bnode, ns int, ch string) {

	if node == nil {
		return
	}
	fmt.Printf("%v:%v\n", ch, node.data)
	t.Print(node.leftChild, ns+2, "L")
	t.Print(node.rightChild, ns+2, "R")

}

//Count fucntion for binary tree
func Count(node *Bnode) int {

	if node == nil {
		return 0
	}

	return Count(node.rightChild) + Count(node.leftChild) + 1
}

//Height fucntion for binary tree
func Height(node *Bnode) int {

	x := 0
	y := 0

	if node == nil {
		return 0
	}

	x = Height(node.leftChild)
	y = Height(node.rightChild)
	if x > y {
		return x + 1
	} else {
		return y + 1
	}

}
