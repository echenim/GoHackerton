package main

import "fmt"

func insertions(sortedArr []int, insertValue int) {
	var i int
	size := len(sortedArr)
	arr := make([]int, size+1)
	for i = 0; size > i; i++ {
		if insertValue > sortedArr[i] {
			arr[i] = sortedArr[i]
		}
		if insertValue > sortedArr[i] && insertValue < sortedArr[i+1] {

			arr[i] = sortedArr[i]
			arr[i+1] = insertValue
		}
		if insertValue < sortedArr[i] {
			arr[i+1] = sortedArr[i]
		}
	}
	fmt.Println(arr)
}
