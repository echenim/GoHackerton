package main

func main() {
	var node *Node

	node = Add(node, "0")
	node = Add(node, "1")
	node = Add(node, "2")
	node = Add(node, "3")

	node.Print()
}
