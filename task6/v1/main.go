package main

import (
	"fmt"
	"time"
)

func runner(stop <-chan bool) {
	for {
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

	time.Sleep(5 * time.Second)
	stop <- true

	time.Sleep(1 * time.Second)
	fmt.Println("exit")
}
