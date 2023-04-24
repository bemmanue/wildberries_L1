package main

import "fmt"

func main() {
	// создаем переменные
	a, b := 0, 100
	fmt.Printf("Начальные значения: a = %d, b = %d\n", a, b)

	// меняем местами значения переменных
	a, b = b, a
	fmt.Printf("Значения после замены: a = %d, b = %d\n", a, b)
}
