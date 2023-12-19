package main

import (
	"fmt"
	"sync"
	"time"
)

func sleep(second int) {
	to := time.After(time.Duration(second) * time.Second)
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		for {
			select {
			case <-to:
				wg.Done()
				return
			}
		}
	}()

	wg.Wait()
}

func main() {
	n := time.Now()
	sleep(10)
	fmt.Println("Hi")
	t := time.Since(n)

	fmt.Println(t)
}
