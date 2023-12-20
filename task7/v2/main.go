package main

import (
	"context"
	"fmt"
	"math/rand"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
	_ = cancel
	// для блокировки одновременной записи и чтение нескольких горутин используем sync.Map
	var sm sync.Map

	for i := 0; i < 5; i++ {
		// запускаем 5 горутин для записи и 5 горутин для чтения из map
		go func() {
			for j := 0; ; j++ {
				source := rand.NewSource(time.Now().UnixNano())
				random := rand.New(source)
				n := random.Intn(10) + 1
				sm.Store(j, fmt.Sprintf("val-%d", n))
				time.Sleep(1 * time.Second)
			}
		}()
		go func(i int) {
			for j := 0; ; j++ {
				val, ok := sm.Load(j)
				if ok {
					fmt.Printf("goroutune-%d - %d:%s\n", i, j, val)
				}
				time.Sleep(1 * time.Second)
			}
		}(i)
	}
	<-ctx.Done()
}
