package main

import "fmt"

func main() {

	H := []int{2, 10, 20, 30, 25, 5, 40, 35}
	fmt.Println("......Original data .........")
	fmt.Println(H)
	for i := 0; i <= len(H)-1; i++ {
		Insert(H, i)
	}
	fmt.Println("......Heap Array formed .........")
	fmt.Println(H)
	for i := len(H) - 1; i > 0; i-- {
		fmt.Println("......deleted data .........")
		k := Delete(H, i)
		fmt.Println(k)
		fmt.Println("......After deleted Heap Array formed .........")
		fmt.Println(H)

	}
	// fmt.Println("......After deleted Heap Array formed .........")
	// fmt.Println(H)
}
