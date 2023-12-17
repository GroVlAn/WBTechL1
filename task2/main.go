package main

import (
	"fmt"
	"sync"
)

func main() {
	// количество горутин работающих с данными из массива
	n := 3
	// искомый слайс значений
	nums := []int{2, 4, 6, 8, 10}
	var wg sync.WaitGroup
	ch := make(chan int, len(nums))

	// обновляем счётчик WaitGroup равный количеству запущенных горутин
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func(i int) {
			for {
				val, open := <-ch

				if !open {
					// если канал закрыт, устанавливаем флаг, что горутина закончила свою работу
					wg.Done()
					return
				}

				//выводим квадарат чисел
				fmt.Println(val * val)
			}
		}(i)
	}

	// наполняем канал значениями
	for _, num := range nums {
		ch <- num
	}

	//закрываем канал, так как добавления значений больше не планируется
	close(ch)
	// ждём пока все горутины закончат работу
	wg.Wait()
}
