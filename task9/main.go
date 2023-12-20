package main

import (
	"fmt"
	"sync"
)

const N = 5

func main() {
	arr := [5]int{1, 6, 3, 53, 5}
	// объявляем два канала
	ch1 := make(chan int, N)
	ch2 := make(chan int, N)
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		// пока второй канал открыт выводим его значения
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
		// пока первый канал открыт пишем из него квадраты значений в канал <-ch2
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

	// Заполняем первый канал значениями из массива
	for i := 0; i < len(arr); i++ {
		ch1 <- arr[i]
	}

	// закрываем первый канал
	close(ch1)
	wg.Wait()
}
