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
func (tree *Btree) Print(b *Bnode, ns int, ch string) {

	if b == nil {
		return
	}
	fmt.Printf("%v:%v\n", ch, b.data)
	tree.Print(b.leftChild, ns+2, "L")
	tree.Print(b.rightChild, ns+2, "R")

}

//Count fucntion for binary tree
func Count(n *Bnode) int {

	if n == nil {
		return 0
	}

	return Count(n.rightChild) + Count(n.leftChild) + 1
}

//Height fucntion for binary tree
func Height(n *Bnode) int {

	x := 0
	y := 0

	if n == nil {
		return 0
	}

	x = Height(n.leftChild)
	y = Height(n.rightChild)
	if x > y {
		return x + 1
	} else {
		return y + 1
	}

}
