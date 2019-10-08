package main

import (
	"unsafe"
)

type stack struct {
	size int
	top  int
	arr  unsafe.Pointer
}
