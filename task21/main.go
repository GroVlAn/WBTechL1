package main

import "fmt"

// CalculatorArea - старый интерфейс, метод кготорого принимает массив из двух элементов
type CalculatorArea interface {
	Area(sides [2]int) int
}

type OldArea struct {
}

// Area - старый метод нахождения площади
func (oa *OldArea) Area(sides [2]int) int {
	return sides[0] * sides[1]
}

// NewArea - новая стркутура, которая имеет две стороны x и y
type NewArea struct {
	x int
	y int
}

// AreaAdapter - структура адаптер, в которую встроена структура NewArea
type AreaAdapter struct {
	NewArea
}

// Area - метод адаптера, реализующий интерфейс CalculatorArea, но использующая в качестве значений
// значения из структуры NewArea
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
