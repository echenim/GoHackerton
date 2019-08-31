package main

import "fmt"

func main() {

	arr := []int{3, 5, 7, 4, 8, 6, 9}
	//arr := []int{}
	first := Add(arr)
	second := first
	fmt.Printf("\nis linkedlist empty %v", first.IsEmpty())
	fmt.Printf("\nHead node = %v", first.Head())
	fmt.Println("")
	first.Print()
	fmt.Println("\n--------------------")
	second = RemoveAtFirstNode(second)
	fmt.Printf("\nis linkedlist empty %v", second.IsEmpty())
	fmt.Printf("\nHead node = %v", second.Head())
	fmt.Println("")
	second.Print()

}
