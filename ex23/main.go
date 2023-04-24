package main

import "fmt"

func remove(arr []int, i int) []int {
	return append(arr[:i], arr[i+1:]...)
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	i := 2

	fmt.Println(remove(arr, i))
}
