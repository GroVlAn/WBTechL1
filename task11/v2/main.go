package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	n1 := 40
	n2 := 40
	nums1 := make([]int, 0, n1)
	nums2 := make([]int, 0, n2)
	union := make([]int, 0)

	for i := 0; i < n1; i++ {
		source := rand.NewSource(time.Now().UnixNano())
		random := rand.New(source)
		nums1 = append(nums1, random.Intn(400)+1)
	}

	for i := 0; i < n2; i++ {
		source := rand.NewSource(time.Now().UnixNano())
		random := rand.New(source)
		nums2 = append(nums2, random.Intn(400)+1)
	}

	for _, n1 := range nums1 {
		for _, n2 := range nums2 {
			if n1 == n2 {
				union = append(union, n1)
			}
		}
	}

	fmt.Println(nums1)
	fmt.Println(nums2)
	fmt.Println(union)
}
