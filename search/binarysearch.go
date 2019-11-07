package main

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
