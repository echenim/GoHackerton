package main

//Queue structure defination
type Queue struct {
	size  int
	front int
	rear  int
	arr   []string
}

//Enqueue insert into the queue
func (q *Queue) Enqueue(s string) {
	if q.rear < q.size-1 {
		q.rear++
		q.arr[q.rear] = s
	}
}

//Dequeue delete from the queue
func Dequeue() {

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
