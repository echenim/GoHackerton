package main

import "fmt"

func factorial(num int) int {
	if num == 0 {
		return 1
	}
	return num * factorial(num-1)
}

func fibinacciInteration(lim int, next int, first int, second int) {
	var i int

	for i = 0; i < lim; i++ {

		if i <= 1 {
			next = i
		} else {
			next = first + second
			first = second
			second = next
		}
		fmt.Println(next)
	}

}
