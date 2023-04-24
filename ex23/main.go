package main

import "fmt"

// remove удаляет i-ый элемент среза
func remove(arr []int, i int) []int {
	// обрабатываем случай, когда исходный срез пустой или индекс выходит за пределы длины
	if len(arr) == 0 || i >= len(arr) {
		return arr
	}

	// создаем новый срез и копируем в него исходный, чтобы у каждого была своя область памяти
	arr2 := make([]int, len(arr))
	copy(arr2, arr)

	return append(arr2[:i], arr2[i+1:]...)
}

func main() {
	arr := []int{1, 2, 3, 4, 5} // срез
	i := 2                      // индекс элемента, который нужно удалить из среза

	fmt.Println(remove(arr, i)) // печатаем срез с удаленным элементом
	fmt.Println(arr)            // изменения не отразились на исходном срезе
}
