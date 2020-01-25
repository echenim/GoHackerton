package main

import (
	"fmt"
)

func main() {
	fmt.Printf("%08b %d \n", 2, 2)
	k := setBit(2, 3)
	fmt.Printf("%08b %d \n", k, k)
	m := clearBit(k, 3)
	fmt.Printf("%08b %d \n", m, m)

	n := modifyBit(k, 3, 1)
	fmt.Printf("%08b %d \n", n, n)
}

// func main() {

// 	// a := 2
// 	// b := 16
// 	// c := 32
// 	// fmt.Printf("%08b\n", a)
// 	// fmt.Printf("%08b\n", b)
// 	// fmt.Printf("%08b\n", c)

// 	// for i := 0; i <= 30; i++ {
// 	// 	k := math.Pow(3, float64(i))
// 	// 	e := int64(k)
// 	// 	fmt.Printf("%d  |  %064b   \n", e, e)
// 	// }

// 	// c := isPowerOfTwo(9)
// 	// fmt.Println(c)
// 	// var a uint
// 	// var b uint
// 	a := 196
// 	// b = 13
// 	// fmt.Printf("%08b  |  %d\n", a, a)
// 	// fmt.Printf("%08b  |  %d\n", b, b)
// 	fmt.Printf("%08b  |  %d\n", a, a)
// 	a &= a
// 	fmt.Printf("%08b  |  %d\n", a, a)
// 	b := int64(a)
// 	d := strconv.FormatInt(b, 2)
// 	fmt.Println(d)
// 	var s string
// 	for _, d := range d {
// 		if d == 49 {
// 			s = s + string(d-1)
// 		}

// 		if d == 48 {
// 			s = s + string(d+1)
// 		}
// 	}

// 	fmt.Println(s)

// 	k, _ := strconv.Atoi(s)

// 	z := uint(k)

// 	fmt.Printf("00%d \n", z)

// 	// c := a & b
// 	// fmt.Printf("%08b  |  %d\n", c, c)

// 	// d := a | b
// 	// fmt.Printf("%08b  |  %d\n", d, d)

// 	// e := 2 ^ 3
// 	// fmt.Printf("%08b  |  %d\n", e, e)

// }

// func isPowerOfTwo(x int) bool {
// 	k := x & (x - 1)
// 	return k == 0
// }
