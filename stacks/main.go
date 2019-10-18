package main

import "fmt"

//func main() {
//	var st stack
//
//	st.size = 10
//	st.top = -1
//
//	st.arr = make([]int, st.size)
//	gen := datagen(st.size)
//	fmt.Printf("initial top :: %v\n", st.top)
//
//	fmt.Println("------ stack growth and information ---------")
//
//	for i := range gen {
//		value := gen[i]
//		if st.top < st.size {
//			st.top++
//			st.Push(value)
//
//		}
//
//		fmt.Printf("stack size :: %v\n", st.size)
//		fmt.Printf("current top :: %v\n", st.top)
//		fmt.Printf("Is stack empty: %v\n", st.IsEmpty())
//		fmt.Printf("Is stack full: %v\n", st.IsFull())
//		fmt.Printf("stack from F->L :: %v\n", st.arr)
//
//	}
//
//	fmt.Println("\n------ stack pop and information ---------")
//
//	for st.top != -1 {
//		st.top--
//		if st.top>-1{
//			st.Pop()
//			fmt.Printf("stack size :: %v\n", st.size)
//			fmt.Printf("current top :: %v\n", st.top)
//			fmt.Printf("Is stack empty: %v\n", st.IsEmpty())
//			fmt.Printf("Is stack full: %v\n", st.IsFull())
//			fmt.Printf("stack from Last -> First :: %v\n", st.arr)
//		}
//	}
//
//}
//
//func datagen(size int) []int {
//	var i int
//	arr := make([]int, size, size)
//	for i = 0; i < size; i++ {
//		arr[i] = rand.Intn(1000)
//	}
//	fmt.Printf("generated data :: %v\n", arr)
//	return arr
//}



func main(){
	list := &SingleList{}
	list.Push(&Node{data: "10", next: nil})
	list.Push(&Node{data: "15", next: nil})
	list.Push(&Node{data: "8", next: nil})
	list.Push(&Node{data: "3", next: nil})
	//node := list.head
	list.Display()
	fmt.Printf("\ntop & size of stack is : %v",list.top)
	fmt.Print("\n")
	list.Pop()
	list.Display()
}