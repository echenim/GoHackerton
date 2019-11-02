package main

func main() {

	tree := &Btree{}
	tree.Insert(&Bnode{data: "echenim", leftChild: nil, rightChild: nil})
	tree.Insert(&Bnode{data: "william", leftChild: nil, rightChild: nil})
	tree.Insert(&Bnode{data: "chukwuma", leftChild: nil, rightChild: nil})
	tree.Insert(&Bnode{data: "christian", leftChild: nil, rightChild: nil})
	tree.Insert(&Bnode{data: "physicist", leftChild: nil, rightChild: nil})
	tree.Insert(&Bnode{data: "myron", leftChild: nil, rightChild: nil})

}
