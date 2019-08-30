package main

import "fmt"

func main() {

	arr := []int{3, 5, 7, 4, 8, 6, 9}
	//arr := []int{}
	first := Add(arr)

	first.Print()
	fmt.Printf("\nis linkedlist empty %v", first.IsEmpty())
}
