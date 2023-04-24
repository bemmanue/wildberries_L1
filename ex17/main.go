package main

import "fmt"

func binarySearch(arr []int, n int) bool {
	left := 0
	right := len(arr) - 1

	for left <= right {
		mid := left + (right-left)/2

		fmt.Println(arr[mid])
		if arr[mid] == n {
			return true
		} else if arr[mid] < n {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return false
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	n := 1

	ok := binarySearch(arr, n)
	fmt.Printf("Массив содержит число %d: %v\n", n, ok)
}
