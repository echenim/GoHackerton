package main

import "fmt"

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
	}
	if node.data <= b.data {

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

//FindMax function for search the tree
func (t *Btree) FindMax() {
	if t.root == nil {
		return
	} else {
		t.root.max(t.root)
	}
}

func (b *Bnode) max(node *Bnode) {
	if b == nil {
		return
	} else if b.rightChild != nil {
		if node.data < b.rightChild.data {
			b.rightChild.max(b.rightChild)
		}
	} else {
		fmt.Printf("search for maximu value found and is %c\n", node.data)
	}

}

//FindMin function for search the tree
func (t *Btree) FindMin() {
	if t.root == nil {
		return
	} else {
		t.root.min(t.root)
	}
}

func (b *Bnode) min(node *Bnode) {
	if b == nil {
		return
	} else if b.leftChild != nil {
		if node.data > b.leftChild.data {
			b.leftChild.min(b.leftChild)
		}

	} else {
		fmt.Printf("search for minimu value found and is %c\n", node.data)
	}

}

// func (b *Bnode) findManx(node *Bnode) {
// 	if b == nil {
// 		return
// 	}
// 	if b.rightChild != nil {
// 		if b.data > b.rightChild.data {
// 			fmt.Printf("not found",b.)
// 		} else {
// 			b.leftChild.findManx(node)
// 		}
// 	}else if b.rightChild != nil && b.data < b.leftChild.data{

// 	} else {
// 		if b.rightChild == nil {
// 			fmt.Println("match not found")
// 		} else {
// 			if parameters == b.rightChild.data {
// 				fmt.Println("match found")
// 			} else {
// 				b.rightChild.search(parameters)
// 			}
// 		}
// 	}
// }

//Search function for search the tree
func (t *Btree) Search(parameters byte) {
	if t.root == nil {
		return
	} else {
		t.root.search(parameters)
	}
}

func (b *Bnode) search(parameters byte) {
	if b == nil {
		return
	}
	if parameters == b.data {
		fmt.Println("match found")
	} else if parameters < b.data {
		if b.leftChild == nil {
			fmt.Println("match not found")
		} else {
			if parameters == b.leftChild.data {
				fmt.Println("match found")
			} else {
				b.leftChild.search(parameters)
			}
		}
	} else {
		if b.rightChild == nil {
			fmt.Println("match not found")
		} else {
			if parameters == b.rightChild.data {
				fmt.Println("match found")
			} else {
				b.rightChild.search(parameters)
			}
		}
	}
}

//PreOrder fucntion for displaying the tree nodes and branch
func (t *Btree) PreOrder(node *Bnode, ch string) {

	if node == nil {
		return
	}
	fmt.Printf("%v:%c  |   ", ch, node.data)
	t.PreOrder(node.leftChild, "L")
	t.PreOrder(node.rightChild, "R")

}

//PostOrder fucntion for displaying the tree nodes and branch
func (t *Btree) PostOrder(node *Bnode, ch string) {

	if node == nil {
		return
	}

	t.PostOrder(node.leftChild, "L")
	t.PostOrder(node.rightChild, "R")
	fmt.Printf("%v:%c  |   ", ch, node.data)

}

//Count fucntion for binary tree
func Count(node *Bnode) int {

	if node == nil {
		return 0
	}

	return Count(node.rightChild) + Count(node.leftChild) + 1
}

//FindHeight fucntion for binary tree
func FindHeight(node *Bnode) int {

	x := 0
	y := 0

	if node == nil {
		return 0
	}

	x = FindHeight(node.leftChild)
	y = FindHeight(node.rightChild)
	if x > y {
		return x + 1
	} else {
		return y + 1
	}

}
