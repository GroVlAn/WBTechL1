package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	n1 := 40
	n2 := 40
	union := make(map[int]int)
	nums1 := make([]int, 0, n1)
	nums2 := make([]int, 0, n2)

	for i := 0; i < n1; i++ {
		source := rand.NewSource(time.Now().UnixNano())
		random := rand.New(source)
		nums1 = append(nums1, random.Intn(40)+1)
	}

	for i := 0; i < n2; i++ {
		source := rand.NewSource(time.Now().UnixNano())
		random := rand.New(source)
		nums2 = append(nums2, random.Intn(40)+1)
	}

	for n := range nums1 {
		_, ok := union[n]

		if ok {
			union[n]++
		} else {
			union[n] = 1
		}
	}

	for n := range nums2 {
		_, ok := union[n]

		if ok {
			union[n]++
		} else {
			union[n] = 1
		}
	}

	for key, val := range union {
		if val > 1 {
			fmt.Println(key)
		}
	}

}
