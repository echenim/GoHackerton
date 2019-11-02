package main

//Btree defination for single linked list
type Btree struct {
	root   *Bnode
	height int
	level  int
}

//Insert function for purpolate the tree
func (tree *Btree) Insert(node *Bnode) *Btree {
	if tree.root == nil {
		tree.root = node
		tree.height = 1
		tree.level = 0

	} else {
		tree.root.insert(node)
	}
	return tree
}

func (b *Bnode) insert(n *Bnode) {
	if b == nil {
		return
	}if b.leftChild == nil {

	} else if b.rightChild == nil {

	}
}
