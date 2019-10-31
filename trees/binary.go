package main

//TreeList defination for single linked list
type TreeList struct {
	head *Node
}

//Insert function for purpolate the tree
func (t TreeList) Insert(node *Node) {

	if t.head == nil {
		t.head = node
		return
	}

	//currentNode = t.head

}
