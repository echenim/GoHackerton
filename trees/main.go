package main

import "fmt"

func main() {

	tree := Btree{}
	tree.Insert(&Bnode{data: "echenim", leftChild: nil, rightChild: nil})
	tree.Insert(&Bnode{data: "william", leftChild: nil, rightChild: nil})
	tree.Insert(&Bnode{data: "chukwuma", leftChild: nil, rightChild: nil})
	tree.Insert(&Bnode{data: "christian", leftChild: nil, rightChild: nil})
	tree.Insert(&Bnode{data: "physicist", leftChild: nil, rightChild: nil})
	tree.Insert(&Bnode{data: "myron", leftChild: nil, rightChild: nil})
	tree.Print(tree.root, 0, "M")
	fmt.Printf("Number of node or branch in this binary tree :  %v", Count(tree.root))
	//Count(tree.root)

}
