package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

func datagen(size int) []int {
	var i int
	arr := make([]int, size, size)
	for i = 0; i < size; i++ {
		arr[i] = rand.Intn(1000)
	}
	fmt.Printf("generated data :: %v\n", arr)
	return arr
}

func main() {

	var q Queue
	var i int

	q.size = 10

	q.arr = make([]string, q.size, q.size)
	fmt.Printf("generated data :: %v\n", q.arr)
	for i = 0; i < q.size; i++ {
		val := strconv.Itoa(rand.Intn(1000))
		q.Enqueue(val)
		fmt.Printf("student queue by number :: %v\n", q.arr)
	}

}
