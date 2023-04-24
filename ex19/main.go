package main

import (
	"fmt"
)

func reverse(s string) string {
	r := []rune(s)

	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}

	return string(r)
}

func main() {
	s := "главрыба"

	fmt.Printf("%s – %s\n", s, reverse(s))
}
