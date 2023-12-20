package main

import (
	"context"
	"fmt"
	"os/signal"
	"syscall"
)

func main() {
	// останавливаем горутину с помощью системных сигналов, используя signal.NotifyContext
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)

	_ = cancel

	ch := make(chan int, 6)

	go func() {
		for val := range ch {
			fmt.Println(val)
		}
		fmt.Println("goroutine is stopped")
	}()

	ch <- 1
	ch <- 2
	ch <- 3
	close(ch)
	<-ctx.Done()
}
