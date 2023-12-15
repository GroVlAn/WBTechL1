package main

import (
	"context"
	"fmt"
	"math/rand"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
	"time"
)

type SyncMap struct {
	mx sync.RWMutex
	m  map[string]string
}

func NewSyncMap() *SyncMap {
	return &SyncMap{
		m: make(map[string]string),
	}
}

func (sm *SyncMap) Set(key, val string) {
	sm.mx.Lock()
	defer sm.mx.Unlock()

	sm.m[key] = val
}

func (sm *SyncMap) Value(key string) (string, string) {
	sm.mx.RLock()
	defer sm.mx.RUnlock()

	return key, sm.m[key]
}

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
	_ = cancel
	syncMap := NewSyncMap()
	for i := 0; i < 5; i++ {
		go func(i int) {
			for j := 0; ; j++ {
				source := rand.NewSource(time.Now().UnixNano())
				random := rand.New(source)
				m := random.Intn(10) + 1
				syncMap.Set(strconv.Itoa(j), fmt.Sprintf("var-%d", m))
				fmt.Printf("goroutine number %d set data: %s:%s \n",
					i, strconv.Itoa(j), fmt.Sprintf("var-%d", m))
				time.Sleep(1 * time.Second)
			}
		}(i)

		go func(i int) {
			for j := 0; ; j++ {
				select {
				case <-time.Tick(1 * time.Second):
					k, v := syncMap.Value(strconv.Itoa(j))
					fmt.Printf("\n\tresult goroutine №%d: %s: %s\n", i, k, v)
				}
			}
		}(i)
	}
	<-ctx.Done()
}
