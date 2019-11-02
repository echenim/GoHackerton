package main

func main() {

	tree := &Btree{}
	tree.Insert(&Bnode{data: 100, leftChild: nil, rightChild: nil})
	tree.Insert(&Bnode{data: -20, leftChild: nil, rightChild: nil})
	tree.Insert(&Bnode{data: -50, leftChild: nil, rightChild: nil})
	tree.Insert(&Bnode{data: -15, leftChild: nil, rightChild: nil})
	tree.Insert(&Bnode{data: -60, leftChild: nil, rightChild: nil})
	tree.Insert(&Bnode{data: 50, leftChild: nil, rightChild: nil})
	tree.Insert(&Bnode{data: 60, leftChild: nil, rightChild: nil})
	tree.Insert(&Bnode{data: 55, leftChild: nil, rightChild: nil})
	tree.Insert(&Bnode{data: 85, leftChild: nil, rightChild: nil})
	tree.Insert(&Bnode{data: 15, leftChild: nil, rightChild: nil})
	tree.Insert(&Bnode{data: 5, leftChild: nil, rightChild: nil})
	tree.Insert(&Bnode{data: -10, leftChild: nil, rightChild: nil})

}
