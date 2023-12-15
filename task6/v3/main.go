package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	ctx, cansel := context.WithCancel(context.Background())
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)

	go func() {
		for {
			select {
			case <-quit:
				fmt.Println("exit")
				cansel()
				return
			case <-time.Tick(1 * time.Second):
				fmt.Println("working...")
			}
		}
	}()

	<-ctx.Done()
}
