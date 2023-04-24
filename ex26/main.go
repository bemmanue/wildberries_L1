package main

import (
	"fmt"
	"unicode"
)

func unique(s string) bool {
	symbols := make(map[byte]bool)

	for _, a := range s {
		a = unicode.ToUpper(a)
		if _, ok := symbols[byte(a)]; ok == true {
			return false
		}
		symbols[byte(a)] = true
	}
	return true
}

func main() {
	arr := []string{"abcd", "abCdefAaf", "aabcd"}

	for _, s := range arr {
		fmt.Printf("Все символы в строке %s уникальные: %v\n", s, unique(s))
	}
}
