package main

import "fmt"

func setBit(n int64, pos int) int64 {
	n |= 1 << pos // | - bitwise OR

	return n
}

func clearBit(n int64, pos int) int64 {
	var mask int64

	mask = ^(1 << pos) // ^ - bitwise NOT
	n &= mask          // & - bitwise AND

	return n
}

func main() {
	var n int64

	n = setBit(n, 7)
	fmt.Println(n)

	n = clearBit(n, 7)
	fmt.Println(n)
}
