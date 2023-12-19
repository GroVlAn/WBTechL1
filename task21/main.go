package main

import "fmt"

type CalculatorArea interface {
	Area(sides [2]int) int
}

type OldArea struct {
}

func (oa *OldArea) Area(sides [2]int) int {
	return sides[0] * sides[1]
}

type NewArea struct {
	x int
	y int
}

type AreaAdapter struct {
	NewArea
}

func (aa *AreaAdapter) Area(sides [2]int) int {
	aa.x = sides[0]
	aa.y = sides[1]

	return aa.x * aa.y
}

func main() {
	calc := CalculatorArea(&AreaAdapter{})
	sides := [2]int{4, 8}
	area := calc.Area(sides)
	fmt.Println(area)
}
