package main

import (
	"fmt"
	"sync"
)

var sum int

func plus(n int, mt *sync.Mutex) {
	mt.Lock()
	defer mt.Unlock()
	sum += n
}

func main() {
	nums := []int{2, 4, 6, 8, 10}
	n := len(nums) * 4 / 10
	var wg sync.WaitGroup
	var mt sync.Mutex
	ch := make(chan int, len(nums))
	fmt.Println(n)
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			for {
				val, open := <-ch

				if !open {
					wg.Done()
					return
				}

				plus(val, &mt)
			}
		}()
	}

	for _, num := range nums {
		ch <- num * num
	}

	close(ch)
	wg.Wait()
	fmt.Println(sum)
}
