package main

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
		// tree.height = 1
		// tree.level = 0

	} else {
		tree.root.insert(node)
	}
	return tree
}

func (b *Bnode) insert(n *Bnode) {
	if b == nil {
		return
	} else if n.data >= b.data {
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
