package main

import (
	"fmt"
	"math"
)

func main() {
	var groups = make(map[int][]float64)
	temperatures := []float64{-31.0, -43.3, -33.3, -49.1, -25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	var key int
	for _, t := range temperatures {
		// вычисляем основание с шагом в 10
		key = int(t - math.Mod(t, 10))

		_, exist := groups[key]

		// если уже есть группа с данным основание, добавляем новый эллемент, если нет создаём новый слайс
		if exist {
			groups[key] = append(groups[key], t)
		} else {
			groups[key] = make([]float64, 0)
			groups[key] = append(groups[key], t)
		}
	}

	fmt.Println(groups)
}
