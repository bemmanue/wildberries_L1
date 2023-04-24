package main

import "fmt"

func main() {
	a, b := 0, 100
	fmt.Printf("Начальные значения: a = %d, b = %d\n", a, b)

	a, b = b, a
	fmt.Printf("Значения после замены: a = %d, b = %d\n", a, b)
}
