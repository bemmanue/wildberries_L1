package main

import "fmt"

// remove удаляет i-ый элемент среза
func remove(arr []int, i int) []int {
	return append(arr[:i], arr[i+1:]...)
}

func main() {
	arr := []int{1, 2, 3, 4, 5} // срез
	i := 2                      // индекс элемента, который нужно удалить из среза

	fmt.Println(remove(arr, i))
}
