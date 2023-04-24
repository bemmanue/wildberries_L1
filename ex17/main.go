package main

import "fmt"

// binarySearch осуществляет поиск значения в отсортированном срезе
func binarySearch(arr []int, n int) bool {
	// определяем крайние индексы, в промежутке которых будет осуществляться поиск
	left := 0
	right := len(arr) - 1

	// ищем значение в промежутке
	for left <= right {
		mid := left + (right-left)/2 // находим середину промежутка

		if arr[mid] == n {
			return true // значение найдено
		} else if arr[mid] < n {
			left = mid + 1 // отбрасываем вторую половину
		} else {
			right = mid - 1 // отбрасываем первую половину
		}
	}

	return false // значение не найдено
}

func main() {
	// отсортированный срез
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}

	// значение, которое нужно найти
	n := 1

	// печатаем результат поиска
	fmt.Printf("Массив содержит число %d: %v\n", n, binarySearch(arr, n))
}
