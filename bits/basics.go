package main

func setBit(x int, position uint) int {
	mask := 1 << position
	return x | mask
}

func clearBit(x int, position uint) int {
	mask := 1 << position
	return x &^ mask
}

func modifyBit(x int, position uint, state int) int {
	mask := 1 << position
	return (x &^ mask) | (-state & mask)
}

func flipBit(x int, position uint) int {
	mask := 1 << position
	return x ^ mask
}

func isBitSet(x int, position uint) bool {
	x >>= position
	return (x & 1) != 0
}
