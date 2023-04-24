package main

import (
	"fmt"
)

// min возвращает меньшее число
func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

// intersect получает два множества в виде map и возвращает пересечение этих множеств в виде map
func intersect(s1, s2 map[int]int) map[int]int {
	s := make(map[int]int)

	for key, v1 := range s1 {
		if v2, ok := s2[key]; ok == true {
			s[key] = min(v1, v2) // добавляем в пересечение количество элементов, которое есть как в первом, так и втором множестве
		}
	}

	return s
}

func main() {
	// создаем первое множество
	s1 := map[int]int{
		4: 1,
		5: 2,
		2: 3,
	}

	// создаем второе множество
	s2 := map[int]int{
		7: 2,
		3: 1,
		5: 1,
	}

	// печатаем результат пересения множеств
	fmt.Println(intersect(s1, s2))
}
