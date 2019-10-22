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
	q.size = 10
	q.front = -1
	q.rear = q.front

	q.arr = make([]string, q.size, q.size)
	fmt.Printf("generated data :: %v\n", q.arr)
	for q.rear < q.size-1 {
		val := strconv.Itoa(rand.Intn(1000))
		q.Enqueue(val)
		fmt.Printf("queue rear :: %v\n", q.front)
		fmt.Printf("queue rear :: %v\n", q.rear)
		fmt.Printf("queue size :: %v\n", q.size)
		fmt.Printf("student queue by number :: %v\n", q.arr)
	}

}
