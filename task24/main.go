package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func NewPoint(x, y float64) *Point {
	return &Point{
		x: x,
		y: y,
	}
}

func (p1 *Point) Distance(p2 *Point) float64 {
	dx := p2.x - p1.x
	dy := p2.y - p1.y

	return math.Sqrt(dx*dx + dy*dy)
}

func main() {
	point1 := NewPoint(5.2, 7.6)
	point2 := NewPoint(19.3, 35.8)

	fmt.Println(point1.Distance(point2))
}
