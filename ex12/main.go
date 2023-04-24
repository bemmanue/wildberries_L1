package main

import "fmt"

// createSet создает множество из последовательности строк
func createSet(arr []string) map[string]bool {
	set := make(map[string]bool)

	for _, val := range arr {
		set[val] = true
	}

	return set
}

func main() {
	arr := []string{"cat", "cat", "dog", "cat", "tree"}

	// печатаем созданное из последовательности строк множество
	fmt.Println(createSet(arr))
}
