package main

import (
	"fmt"
	"sync"
)

const N = 5

func main() {
	arr := [5]int{1, 6, 3, 53, 5}
	ch1 := make(chan int, N)
	ch2 := make(chan int, N)
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		for {
			val, open := <-ch2

			if !open {
				wg.Done()
				return
			}

			fmt.Println(val)
		}
	}()

	wg.Add(1)
	go func() {
		for {
			val, open := <-ch1

			if !open {
				wg.Done()
				close(ch2)
				return
			}

			ch2 <- val * val
		}
	}()

	for i := 0; i < len(arr); i++ {
		ch1 <- arr[i]
	}

	close(ch1)
	wg.Wait()
}
