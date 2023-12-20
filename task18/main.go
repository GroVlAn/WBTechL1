package main

import (
	"fmt"
	"sync"
)

// Increment - структура счётчик
// mt - мьютекс для ограничения увеличения счётчика из нескольких горутин
// inc - счётчик типа int
type Increment struct {
	mt  sync.Mutex
	inc int
}

// Increase - метод для увеличения счётчика
func (i *Increment) Increase() {
	i.mt.Lock()
	defer i.mt.Unlock()
	i.inc++
}

// Inc - метод для получения значения счётчика
func (i *Increment) Inc() int {
	return i.inc
}

func main() {
	inc := new(Increment)
	n := 9999
	var wg sync.WaitGroup

	// запусчкаем n горутин, в которых будет увеличиваться счётчик
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			for i := 0; i < 10; i++ {
				inc.Increase()
			}
		}()
	}

	wg.Wait()
	fmt.Println(inc.Inc())

}
