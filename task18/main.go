package main

import (
	"fmt"
	"sync"
)

type Increment struct {
	mt  sync.Mutex
	inc int
}

func (i *Increment) Increase() {
	i.mt.Lock()
	defer i.mt.Unlock()
	i.inc++
}

func (i *Increment) Inc() int {
	return i.inc
}

func main() {
	inc := new(Increment)
	n := 9999
	var wg sync.WaitGroup

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
