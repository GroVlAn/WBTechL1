package main

import (
	"fmt"
	"time"
)

const N = 5 * time.Second

func main() {
	// Ставим таймер, который вернёт значение в канал по истечению времени
	to := time.After(N)
	ch := make(chan int)
	quit := make(chan bool)

	go func() {
		defer close(ch)
		for i := 1; ; i++ {
			// ждём пока в канал <-to поступит значение, а пока записываем в канал <-ch
			select {
			case <-to:
				// разблокируем канал <-quit для выходи из программы и выодим из горутины
				quit <- true
				break
			default:
				ch <- i
				time.Sleep(1 * time.Second)
			}

		}
	}()

	for {
		// пока данные из канала <-ch поступают, выводим значения, если данные поступили в <-quit выходим из программы
		select {
		case val, open := <-ch:
			if !open {
				return
			}
			fmt.Println(val)
		case <-quit:
			fmt.Println("exit")
			return
		}
	}
}
