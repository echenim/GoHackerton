package main

import "fmt"

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

	tree.Print(tree.root, 0, "M")
	fmt.Printf("Number of node or branch in this binary tree :  %v\n", Count(tree.root))
	fmt.Printf("Height of this binary tree :  %v\n", Height(tree.root))
	//Count(tree.root)

}
