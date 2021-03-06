package main

import "fmt"

func main() {

	tree := Btree{}
	tree.Insert(&Bnode{data: 'F', leftChild: nil, rightChild: nil})
	tree.Insert(&Bnode{data: 'B', leftChild: nil, rightChild: nil})
	tree.Insert(&Bnode{data: 'A', leftChild: nil, rightChild: nil})
	tree.Insert(&Bnode{data: 'D', leftChild: nil, rightChild: nil})
	tree.Insert(&Bnode{data: 'C', leftChild: nil, rightChild: nil})
	tree.Insert(&Bnode{data: 'E', leftChild: nil, rightChild: nil})

	tree.Insert(&Bnode{data: 'G', leftChild: nil, rightChild: nil})
	// tree.Insert(&Bnode{data: 'I', leftChild: nil, rightChild: nil})
	// tree.Insert(&Bnode{data: 'H', leftChild: nil, rightChild: nil})

	// tree.Insert(&Bnode{data: 'J', leftChild: nil, rightChild: nil})
	// tree.Insert(&Bnode{data: 'L', leftChild: nil, rightChild: nil})
	// tree.Insert(&Bnode{data: 'K', leftChild: nil, rightChild: nil})

	// tree.Insert(&Bnode{data: 'M', leftChild: nil, rightChild: nil})
	// tree.Insert(&Bnode{data: 'O', leftChild: nil, rightChild: nil})
	// tree.Insert(&Bnode{data: 'N', leftChild: nil, rightChild: nil})

	// tree.Insert(&Bnode{data: 'P', leftChild: nil, rightChild: nil})
	// tree.Insert(&Bnode{data: 'R', leftChild: nil, rightChild: nil})
	// tree.Insert(&Bnode{data: 'Q', leftChild: nil, rightChild: nil})

	//tree.Insert(&Bnode{data: 'J', leftChild: nil, rightChild: nil})

	fmt.Print("Pre Order tree  :  ")
	tree.PreOrder(tree.root, "M")
	fmt.Println()
	// fmt.Print("Post Order tree :  ")
	// tree.PostOrder(tree.root, "M")

	countNumberOfNode := Count(tree.root)
	fmt.Printf("\nNumber of node or branch in this binary tree :  %v", countNumberOfNode)
	treeheight := FindHeight(tree.root)
	fmt.Printf("\nHeight of this binary tree :  %v", treeheight)

	fmt.Print("\nSearch  : ")
	tree.Search('G')
	tree.FindMax()
	tree.FindMin()
	balanceFactor := BalanceFactor(tree.root.rightChild)
	fmt.Printf("\nBalance Factor  : %v", balanceFactor)

	// fmt.Printf("Number of node or branch in this binary tree :  %v\n", Count(tree.root))
	// fmt.Printf("Height of this binary tree :  %v\n", Height(tree.root))
	//fmt.Printf("%v", Search(tree.root, 35).data)

}
