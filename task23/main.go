package main

import "fmt"

func main() {
	arr := []int{12, 421, 41, 4515, -12, 412, 4521, -24}
	indexDel := 3

	fmt.Println(arr)
	arr = append(arr[:indexDel], arr[indexDel+1:]...)

	fmt.Println(arr)
}
