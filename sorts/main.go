package main

import "fmt"

func main() {
	//unsorted := []int{10, 8, 1, 88, 64, 21, 19, 4, 3, 2}

	// select sort
	// result := selectSort(unsorted)
	// fmt.Println(result)

	//insertion sort
	// insertions(result, 14)
	// fmt.Println(result)

	//bubble sort
	//bubbleSort(unsorted)
	//fmt.Println(unsorted)

	//quick sort
	// quickSort(unsorted, 0, 9)
	// fmt.Println(unsorted)

	//merge sort
	sortedA := []int{1, 3, 5, 7, 8}
	sortedB := []int{2, 4, 6, 9, 10, 11, 12}
	rs := merger(sortedA, sortedB, len(sortedA), len(sortedB))
	fmt.Println(rs)

}
