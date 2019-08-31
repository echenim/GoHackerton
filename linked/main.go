package main

import "fmt"

func main() {

	arr := []int{3, 5, 7, 4, 8, 6, 9}
	//arr := []int{}
	next := Add(arr)
	// first := Add(arr)
	// fmt.Printf("\nis linkedlist empty %v", first.IsEmpty())
	// fmt.Printf("\nHead node = %v", first.Head())
	// fmt.Printf("\nlinkedlist has size = %v", first.Size())
	// fmt.Println("")
	// first.Print()
	// next := first
	for next.next != nil {

		fmt.Println("\n--------------------")
		next.RemoveAtFirstNode()
		fmt.Printf("\nis linkedlist empty %v", next.IsEmpty())
		fmt.Printf("\nHead node = %v", next.Head())
		fmt.Printf("\nlinkedlist has size = %v", next.Size())
		fmt.Println("")
		next.Print()
		next = next.next
	}

}
