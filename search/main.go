package main

func main() {

	tree := Btree{}
	tree.Insert(&Bnode{data: 50, leftChild: nil, rightChild: nil})
	tree.Insert(&Bnode{data: 30, leftChild: nil, rightChild: nil})
	tree.Insert(&Bnode{data: 65, leftChild: nil, rightChild: nil})
	tree.Insert(&Bnode{data: 40, leftChild: nil, rightChild: nil})
	tree.Insert(&Bnode{data: 25, leftChild: nil, rightChild: nil})
	tree.Insert(&Bnode{data: 55, leftChild: nil, rightChild: nil})

	tree.Insert(&Bnode{data: 75, leftChild: nil, rightChild: nil})
	tree.Insert(&Bnode{data: 35, leftChild: nil, rightChild: nil})
	tree.Insert(&Bnode{data: 45, leftChild: nil, rightChild: nil})

	tree.Insert(&Bnode{data: 60, leftChild: nil, rightChild: nil})

}
