package main

import (
	"fmt"
	"unicode"
)

// unique проверяет символы строки на уникальность
func unique(s string) bool {
	symbols := make(map[byte]bool)

	for _, a := range s {
		// приводим символ к единому регистру, поскольку функция проверки должна быть регистронезависимой
		a = unicode.ToUpper(a)

		// проверяем, нет ли в мапе такого символа
		if _, ok := symbols[byte(a)]; ok == true {
			return false
		}

		symbols[byte(a)] = true // записываем в мапу символ
	}

	return true
}

func main() {
	arr := []string{"abcd", "abCdefAaf", "aabcd"}

	for _, s := range arr {
		fmt.Printf("Все символы в строке %s уникальные: %v\n", s, unique(s))
	}
}
