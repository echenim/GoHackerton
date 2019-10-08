package main

type stack struct {
	size int
	top  int
	arr  []int
}

func (s stack) Push(a []int) stack {
	var i int
	s.top++
	for i = 0; i < s.size; i++ {
		s.arr[i] = a[i]
		s.top++
	}
	return s
}
