package main

import (
	"fmt"
)

// quickSort возвращает сортированный срез. Производит быструю сортировку за O(n^2)
func quickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	point := arr[0] // опорная точка

	var less []int    // для значений, которые меньше значения опорной точки
	var greater []int // для значений, которые больше значения опорной точки

	// в цикле проходим по всем элементам, кроме первого и сортируем значения на две группы: less и greater
	for i := range arr[1:] {
		if arr[i] <= point {
			less = append(less, arr[i])
		} else {
			greater = append(greater, arr[i])
		}
	}

	// с помощью рекурсии вычисляем подмассивы и складываем результат
	arr = append(quickSort(less), point)
	arr = append(arr, quickSort(greater)...)

	return arr
}

func main() {
	arr := []int{11, 2, 93, 134, 5, 60, 17, 28, 9, 318, 199, 20}

	fmt.Println(quickSort(arr))
}
