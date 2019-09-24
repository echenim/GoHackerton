package main

import (
	"unsafe"
)

type stack struct {
	size  int
	top   int
	array unsafe.Pointer
}

//Push fi le in the stack
func (s stack) Push() {

}
