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
	// создаём контекст и инициализируем базовый контекст передав context.Background
	ctx, cancel := context.WithCancel(context.Background())
	quit := make(chan os.Signal)
	// ожидаем передаю системных сигналов и передаём их в канал quit
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
	ch := make(chan int)

	workersCount := 1
	args := os.Args

	go func() {
		// в случае если поступил сикгнал о выходе из программы горутина раблокируется данная горутина
		// и будет вызвана функция cancel передав данные в канал ctx.Done()
		<-quit
		cancel()
		fmt.Println("exit")
	}()

	// проверяем что в переданных аргументах есть данные
	if len(args) > 1 {
		numWorkers, err := strconv.Atoi(args[1])
		// проверяем, что первым аргументом было переданно число, если это так, то это число воркеров
		// которое нужно запустить
		if err == nil {
			workersCount = numWorkers
		}
	}

	go func() {
		defer close(ch)
		for {
			// бесконечно генерим значения, которые отправляем в канал
			source := rand.NewSource(time.Now().UnixNano())
			random := rand.New(source)
			ch <- random.Intn(10) + 1
		}
	}()

	// создаём воркеры, которые будут бесконечно читать данные из канала, пока программу не остановят
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

	// блокируем основную горутину, ожидая пока в канал контектса поступят данные
	<-ctx.Done()
}
