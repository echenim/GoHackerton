package main

import "fmt"

func main() {

	H := []int{0, 2, 5, 8, 9, 4, 10, 7}
	fmt.Println(H)
	for i := 2; i <= len(H)-1; i++ {
		Insert(H, i)
		fmt.Println(H)
	}

}
