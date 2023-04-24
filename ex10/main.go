package main

import (
	"fmt"
)

func sort(arr []float64) map[int][]float64 {
	data := make(map[int][]float64)

	for _, n := range arr {
		group := int(n) / 10 * 10
		data[group] = append(data[group], n)
	}

	return data
}

func main() {
	arr := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	fmt.Println(sort(arr))
}
