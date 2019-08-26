package main

func main() {
	var node *Node

	node = add(node, "0")
	node = add(node, "1")
	node = add(node, "2")
	node = add(node, "3")

	node.print()
}
