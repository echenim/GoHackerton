package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var st stack
	st.size = 10
	st.top = -1

	st.arr = make([]int, st.size)
	gen := datagen(st.size)
	fmt.Printf("initial top :: %v\n", st.top)
	k := st.Push(gen)
	fmt.Printf("stack size :: %v\n", k.size)
	fmt.Printf("current top :: %v\n", k.top)
	fmt.Printf("stack from F->L :: %v", k.arr)

}

func datagen(size int) []int {
	var i int
	arr := make([]int, size, size)
	for i = 0; i < size; i++ {
		arr[i] = rand.Intn(1000)
	}
	fmt.Printf("generated data :: %v\n", arr)
	return arr
}
