package main

//DoubleLinkedlist linked list
type DoubleLinkedlist struct {
	head *Dnode
}

func append(arr []string) *Dnode {

	var first *Dnode = new(Dnode) //there will be a runtime error if the point *Node is not initialize as new(Node)
	var last *Dnode = new(Dnode)  //there will be a runtime error if the point *Node is not initialize as new(Node)

	var i int
	for i = 0; i < len(arr); i++ {
		if i == 0 {
			first.prev = nil
			first.data = arr[i]
			first.next = nil
			last = first
		}

		if i > 0 {
			var t *Dnode = new(Dnode) //there will be a runtime error if the point *Node is not initialize as new(Node)
			t.data = arr[i]
			t.prev = last
			t.next = last.next
			last.next = t
			last = t
		}
	}

	return last
}
