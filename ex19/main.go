package main

import (
	"fmt"
)

// reverse переворачивает символы в строке
func reverse(s string) string {
	r := []rune(s) // создаем срез рун для корректного отображения unicode символов

	// меняем местами элементы среза
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}

	return string(r)
}

func main() {
	s := "главрыба"

	// печатаем результат reverse функции
	fmt.Printf("%s – %s\n", s, reverse(s))
}
