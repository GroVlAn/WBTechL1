package main

import (
	"fmt"
	"time"
)

const N = 5 * time.Second

func main() {
	to := time.After(N)
	ch := make(chan int)
	quit := make(chan bool)

	go func() {
		defer close(ch)
		for i := 1; ; i++ {
			select {
			case <-to:
				quit <- true
				break
			default:
				ch <- i
				time.Sleep(1 * time.Second)
			}

		}
	}()

	for {
		select {
		case val, open := <-ch:
			if !open {
				return
			}
			fmt.Println(val)
		case <-quit:
			return
		}
	}
}
