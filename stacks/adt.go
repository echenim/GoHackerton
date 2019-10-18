package main

type stack struct {
	size int
	top  int
	arr  []int
}

//Push insert into stack value
func (s stack) Push(a int) {
	s.arr[s.top] = a
}

func (s stack) Pop() {
	s.arr[s.top] = 0
}

// func (s stack) Peek() {

// }

// func (s stack) StackTop() {

// }

func (s stack) IsEmpty() bool {
	status := false
	if s.top == -1 {
		status = true
	}
	return status
}

func (s stack) IsFull() bool {
	status := false
	if s.top == s.size {
		status = true
	}
	return status
}
