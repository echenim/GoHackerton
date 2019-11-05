package main

import (
	"fmt"
	_ "image/jpeg"
)

//Btree defination for single linked list
type Btree struct {
	root *Bnode
	// height int
	// level  int
}

//Insert function for purpolate the tree
func (tree *Btree) Insert(node *Bnode) *Btree {
	if tree.root == nil {
		tree.root = node
	} else {
		tree.root.insert(node)
	}
	return tree
}

func (b *Bnode) insert(n *Bnode) {
	if b == nil {
		return
	} else if b.leftChild == nil {
		if b.leftChild == nil {
			b.leftChild = n
		} else {
			b.leftChild.insert(n)
		}
	} else {
		if b.rightChild == nil {
			b.rightChild = n
		} else {
			b.rightChild.insert(n)
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
