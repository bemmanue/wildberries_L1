package main

import (
	"fmt"
)

func quickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	var less []int
	var greater []int

	point := arr[0]

	for i := range arr[1:] {
		if arr[i] <= point {
			less = append(less, arr[i])
		} else {
			greater = append(greater, arr[i])
		}
	}

	arr = append(quickSort(less), point)
	arr = append(arr, quickSort(greater)...)

	return arr
}

func main() {
	arr := []int{11, 2, 93, 134, 5, 60, 17, 28, 9, 318, 199, 20}

	fmt.Println(quickSort(arr))
}
