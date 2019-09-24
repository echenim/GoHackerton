package main

import "fmt"

func main() {

	i := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("original data ::  %v \n", i)
	ReverseInts(i)
	fmt.Printf("reversed data ::  %v \n", i)
	s := []string{"boy", "girl", "man", "woman"}
	fmt.Printf("original data ::  %v \n", s)
	ReverseString(s)
	fmt.Printf("reversed data ::  %v \n", s)

}
