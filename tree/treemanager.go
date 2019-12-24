package main

import (
	"fmt"
	"math/rand"
)

// A Tree is a binary tree with integer values.
type Tree struct {
	left  *Tree
	data  int
	right *Tree
}

//Insert create tree data
func Insert(n, k int) *Tree {
	var t *Tree
	for _, v := range rand.Perm(n) {
		t = insert(t, (1+v)*k)
	}
	return t
}

func insert(t *Tree, value int) *Tree {
	if t == nil {
		return &Tree{nil, value, nil}
	}

	if value < t.data {
		t.left = insert(t.left, value)
	} else {
		t.right = insert(t.right, value)
	}

	return t
}

//Print fucntion for displaying the tree nodes and branch
func (t *Tree) Print(ch string) {

	if t == nil {
		return
	}
	fmt.Printf("%v:%v\n", ch, t.data)
	t.left.Print("L")
	t.right.Print("R")

}

// Walk traverses a tree depth-first,
// sending each Value on a channel.
func Walk(t *Tree, ch chan int) {
	if t == nil {
		return
	}
	Walk(t.left, ch)
	ch <- t.data
	Walk(t.right, ch)
}

// Walker launches Walk in a new goroutine,
// and returns a read-only channel of values.
func Walker(t *Tree) <-chan int {
	ch := make(chan int)
	go func() {
		Walk(t, ch)
		close(ch)
	}()
	return ch
}

// Compare reads values from two Walkers
// that run simultaneously, and returns true
// if t1 and t2 have the same contents.
func Compare(t1, t2 *Tree) bool {
	c1, c2 := Walker(t1), Walker(t2)
	for {
		v1, ok1 := <-c1
		v2, ok2 := <-c2
		if !ok1 || !ok2 {
			return ok1 == ok2
		}
		if v1 != v2 {
			break
		}
	}
	return false
}
