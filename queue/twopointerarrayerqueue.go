package main

import "fmt"

//Queue structure defination
type Queue struct {
	size  int
	front int
	rear  int
	arr   []string
}

//Enqueue insert into the queue
func (q *Queue) Enqueue(s string) {

	if q.front == q.rear {

		fmt.Print("Queue is empty\n\n")

	} else if q.rear < q.size {

		q.arr[q.rear] = s
		fmt.Printf("queue front :: %v\n", q.front)
		fmt.Printf("queue rear :: %v\n", q.rear)
		fmt.Printf("queue size :: %v\n", q.size)
		fmt.Printf("student queue by number :: %v \n\n", q.arr)

	}

}

//Dequeue delete from the queue
func (q *Queue) Dequeue() {
	q.front++
	if q.front < q.rear {
		q.arr[q.front] = ""
		fmt.Printf("queue front :: %v\n", q.front)
		fmt.Printf("queue rear :: %v\n", q.rear)
		fmt.Printf("queue size :: %v\n", q.size)
		fmt.Printf("Current Queue Data  : %v \n\n", q.arr)
	} else if q.front == q.rear {
		q.arr[q.front] = ""
		fmt.Printf("queue front :: %v\n", q.front)
		fmt.Printf("queue rear :: %v\n", q.rear)
		fmt.Printf("queue size :: %v\n", q.size)
		fmt.Printf("Current Queue is empty \nReason Front pointer is (%v) and Rear Pointer is (%v) both are equal\n", q.front, q.rear)
		// q.front = -1
		// q.rear = q.front
		// fmt.Printf("Front and Rear pointer reset to %v  %v", q.front, q.rear)
	}

}

//IsEmpty queue is empty
func IsEmpty() bool {

	return true
}

//IsFull queue is full
func IsFull() bool {

	return true
}

//First first element in the queue
func First() {

}

//Last last element in the queue
func Last() {

}
