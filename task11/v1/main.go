package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// создаём два слайса вместимости 40
	n1 := 40
	n2 := 40
	union := make(map[int]int)
	nums1 := make([]int, 0, n1)
	nums2 := make([]int, 0, n2)

	// заполняем оба салйса рандомными значениями
	for i := 0; i < n1; i++ {
		source := rand.NewSource(time.Now().UnixNano())
		random := rand.New(source)
		nums1 = append(nums1, random.Intn(400))
	}

	for i := 0; i < n2; i++ {
		source := rand.NewSource(time.Now().UnixNano())
		random := rand.New(source)
		nums2 = append(nums2, random.Intn(400))
	}

	// записываем в map все значения из первого слайса
	for _, n := range nums1 {
		_, ok := union[n]

		if !ok {
			union[n] = 1
		}
	}

	// заполняем map значениями из второго слайса, и если такое значение есть, увеличиваем ещё счётчик на 1
	for n := range nums2 {
		_, ok := union[n]

		if ok {
			union[n]++
		} else {
			union[n] = 1
		}
	}

	// выводим только те значения из map, чей счётчик больше 1
	for key, val := range union {
		if val > 1 {
			fmt.Println(key)
		}
	}

}
