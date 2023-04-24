package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func New(x, y float64) Point {
	return Point{
		x: x,
		y: y,
	}
}

func distance(p1, p2 Point) float64 {
	w := math.Abs(p1.x - p2.x)
	h := math.Abs(p1.y - p2.y)

	return math.Sqrt(math.Pow(w, 2) + math.Pow(h, 2))
}

func main() {
	point1 := New(2, 2)
	point2 := New(4, 4)

	fmt.Printf(
		"Расстояние между точками %v и %v равно %.2f",
		point1,
		point2,
		distance(point1, point2),
	)
}
