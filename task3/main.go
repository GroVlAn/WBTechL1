package main

import (
	"fmt"
	"sync"
)

var sum int

func plus(n int, mt *sync.Mutex) {
	// Используем мютекс для ограничения записи нескольких горутин
	mt.Lock()
	defer mt.Unlock()
	sum += n
}

func main() {
	nums := []int{2, 4, 6, 8, 10}
	// 40% элементов - количество используемых горутин
	n := len(nums) * 4 / 10
	var wg sync.WaitGroup
	var mt sync.Mutex
	ch := make(chan int, len(nums))

	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			for {
				// читаем из канал, пока он открыт, если он закрыт, завершаем горутину
				val, open := <-ch

				if !open {
					wg.Done()
					return
				}

				plus(val, &mt)
			}
		}()
	}

	// записываем в канал массив
	for _, num := range nums {
		ch <- num * num
	}

	// закрываем канал, ждём завершения всех горутин и выводим рещультат
	close(ch)
	wg.Wait()
	fmt.Println(sum)
}
