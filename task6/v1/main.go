package main

import (
	"fmt"
	"time"
)

func runner(stop <-chan bool) {
	for {
		// останавливаем горутину, если в канал <-stop постубити булевское значение
		select {
		default:
			fmt.Println("working...")
			time.Sleep(1 * time.Second)
		case <-stop:
			fmt.Println("stopped")
			return
		}
	}
}

func main() {
	stop := make(chan bool)

	go runner(stop)

	// Через 5 секунд записываем в канал stop - true для остановки горутины
	time.Sleep(5 * time.Second)
	stop <- true

	time.Sleep(1 * time.Second)
	fmt.Println("exit")
}
