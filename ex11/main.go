package main

import (
	"fmt"
)

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func intersect(s1, s2 map[int]int) map[int]int {
	s := make(map[int]int)

	for key, v1 := range s1 {
		if v2, ok := s2[key]; ok == true {
			s[key] = min(v1, v2)
		}
	}

	return s
}

func main() {
	s1 := map[int]int{
		4: 1,
		5: 2,
		2: 3,
	}

	s2 := map[int]int{
		7: 2,
		3: 1,
		5: 1,
	}

	fmt.Println(intersect(s1, s2))
}
