package main

import "fmt"

func main() {
	list := &SingleList{}

	list.Enqueue(&Node{data: "10", next: nil})
	list.Display()
	fmt.Printf("\nfront pointer is %v & rear pointer is : %v", list.front.data, list.rear.data)
	fmt.Print("\n")

	list.Enqueue(&Node{data: "15", next: nil})
	list.Display()
	fmt.Printf("\nfront pointer is %v & rear pointer is : %v", list.front.data, list.rear.data)
	fmt.Print("\n")

	list.Enqueue(&Node{data: "20", next: nil})
	list.Display()
	fmt.Printf("\nfront pointer is %v & rear pointer is : %v", list.front.data, list.rear.data)
	fmt.Print("\n")

	list.Enqueue(&Node{data: "25", next: nil})
	list.Display()
	fmt.Printf("\nfront pointer is %v & rear pointer is : %v", list.front.data, list.rear.data)
	fmt.Print("\n")

	fmt.Print("\n-----------start dequeue------------\n")
	for n := list.head; n != nil; n = n.next {
		list.Dequeue()
		list.Display()
		fmt.Printf("\nfront pointer is %v & rear pointer is : %v", list.front.data, list.rear.data)
		fmt.Print("\n")
	}

}

// func datagen(size int) []int {
// 	var i int
// 	arr := make([]int, size, size)
// 	for i = 0; i < size; i++ {
// 		arr[i] = rand.Intn(1000)
// 	}
// 	fmt.Printf("generated data :: %v\n", arr)
// 	return arr
// }

// func main() {

// 	var q Queue
// 	q.size = 3
// 	q.front = -1
// 	q.rear = q.front
// 	var i int

// 	q.arr = make([]string, q.size, q.size)
// 	fmt.Printf("generated data :: %v\n", q.arr)
// 	fmt.Println("----------------start enqueue results --------------------")

// 	for q.rear != q.size {
// 		val := strconv.Itoa(rand.Intn(1000))
// 		q.Enqueue(val)
// 		if q.rear < q.size {
// 			q.rear++
// 		}
// 	}

// 	for i = 0; i < q.size; i++ {
// 		val := strconv.Itoa(rand.Intn(1000))
// 		q.rear = i
// 		q.Enqueue(val)
// 	}

// 	fmt.Println("----------------end enqueue results --------------------\n----------------start dequeue results --------------------")

// 	for q.front <= q.rear {
// 		q.Dequeue()
// 	}

// }
