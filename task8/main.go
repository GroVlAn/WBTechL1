package main

import (
	"fmt"
	"os"
	"strconv"
)

const N int64 = 42

func main() {
	i := 0
	currentNumber := N

	args := os.Args

	if len(args) > 1 {
		iB, err := strconv.Atoi(args[1])
		if err == nil {
			i = iB
		}
		if i >= 64 {
			i = 0
		}
	}
	if len(args) > 2 {
		newCunN, err := strconv.Atoi(args[2])
		if err == nil {
			currentNumber = int64(newCunN)
		}
	}
	fmt.Println(i)
	fmt.Println(strconv.FormatInt(currentNumber, 2))
	fmt.Println(currentNumber ^ int64(1)<<i)
	fmt.Println(strconv.FormatInt(currentNumber^int64(1)<<i, 2))
}
