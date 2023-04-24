package main

import (
	"fmt"
	"math"
)

// Point хранит координаты точки
type Point struct {
	x float64
	y float64
}

// New конструктор Point
func New(x, y float64) Point {
	return Point{
		x: x,
		y: y,
	}
}

// distance находит расстояние между двумя точками с помощью теоремы Пифагора
func distance(p1, p2 Point) float64 {
	w := math.Abs(p1.x - p2.x) // длина первого катета
	h := math.Abs(p1.y - p2.y) // длина второго катета

	return math.Sqrt(math.Pow(w, 2) + math.Pow(h, 2)) // гипотенуза
}

func main() {
	// создаем две точки
	point1 := New(2, 2)
	point2 := New(4, 4)

	// печатаем расстояние между точками
	fmt.Printf(
		"Расстояние между точками %v и %v равно %.2f",
		point1,
		point2,
		distance(point1, point2),
	)
}
