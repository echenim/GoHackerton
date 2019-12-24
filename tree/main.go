package main

import "fmt"

func main() {

	t1 := Insert(5, 1)
	t2 := Insert(5, 1)
	fmt.Println(Compare(t1, t2), "Same Contents")
	fmt.Println(Compare(t1, t2), "Differing Sizes")
	fmt.Println(Compare(t1, t2), "Differing Values")
	fmt.Println(Compare(t1, t2), "Dissimilar")
	t1.Print("")
	fmt.Println("----------------------------------")
	t2.Print("")

}
