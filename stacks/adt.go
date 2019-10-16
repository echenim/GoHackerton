package main

import "fmt"

type stack struct {
	size int
	top  int
	arr  []int
}

func (s stack) Push(a int) stack {
	if s.top < s.size {
		s.top++
		s.arr[s.top] = a
	}
	return s
}

func (s stack) Pop() {

	if s.top == -1 {
		fmt.Println("stack underflow")
	}

	for s.top != -1 {
		fmt.Printf(" %v", s.arr[s.top])
		s.top--
	}

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
