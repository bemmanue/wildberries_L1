package main

import "fmt"

// setBit устанавливает i-ый бит в 1
func setBit(n int64, i int) int64 {
	n |= 1 << i // | - bitwise OR

	return n
}

// clearBit устанавливает i-ый бит в 0
func clearBit(n int64, i int) int64 {
	var mask int64

	mask = ^(1 << i) // ^ - bitwise NOT
	n &= mask        // & - bitwise AND

	return n
}

func main() {
	var n int64

	n = setBit(n, 7)
	fmt.Println(n)

	n = clearBit(n, 7)
	fmt.Println(n)
}
