package main

import (
	"fmt"
	"strings"
)

// reverse переворачивает слова в строке
func reverse(s string) string {
	words := strings.Split(s, " ") // получаем срез слов

	// меняем местами элементы среза
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}

	return strings.Join(words, " ") // объединяем элементы среза (слова) в одну строку
}

func main() {
	s := "snow dog sun"

	// печатаем результат reverse функции
	fmt.Printf("%s – %s\n", s, reverse(s))
}
