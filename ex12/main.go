package main

import "fmt"

func createSet(arr []string) map[string]int {
	set := make(map[string]int)

	for _, val := range arr {
		set[val]++
	}

	return set
}

func main() {
	arr := []string{"cat", "cat", "dog", "cat", "tree"}

	fmt.Println(createSet(arr))
}
