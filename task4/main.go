package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
	ch := make(chan int)

	workersCount := 1
	args := os.Args

	go func() {
		<-quit
		cancel()
		fmt.Println("exit")
	}()
	if len(args) > 1 {
		numWorkers, err := strconv.Atoi(args[1])

		if err == nil {
			workersCount = numWorkers
		}
	}

	go func() {
		defer close(ch)
		for {
			source := rand.NewSource(time.Now().UnixNano())
			random := rand.New(source)
			ch <- random.Intn(10) + 1
		}
	}()

	for i := 0; i < workersCount; i++ {
		go func() {
			for {
				select {
				case val, open := <-ch:
					if !open {
						fmt.Println("channel is closed")
						return
					}
					fmt.Println(val)
					time.Sleep(1 * time.Second)
				}
			}
		}()
	}

	<-ctx.Done()
}
