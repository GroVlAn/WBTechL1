package main

import (
	"fmt"
	"sync"
)

func main() {
	nums := []int{2, 4, 6, 8, 10}
	var sum int
	var wg sync.WaitGroup

	wg.Add(len(nums))
	for _, num := range nums {
		go func(num int) {
			sum += num * num
			wg.Done()
		}(num)
	}

	wg.Wait()
	fmt.Println(sum)
}
