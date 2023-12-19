package main

import (
	"fmt"
	"math/big"
)

func main() {
	sum := new(big.Int)
	sub := new(big.Int)
	mul := new(big.Int)
	div := new(big.Int)
	num1 := new(big.Int)
	num2 := new(big.Int)

	fmt.Print("Enter number 1: ")
	fmt.Scan(num1)

	fmt.Print("Enter number 2: ")
	fmt.Scan(num2)

	sum.Add(num1, num2)
	sub.Sub(num1, num2)
	mul.Mul(num1, num2)

	if num2.Cmp(big.NewInt(0)) != 0 {
		div.Div(num1, num2)
	} else {
		fmt.Println("number 2 equal 0")
	}

	fmt.Println(sum)
	fmt.Println(sub)
	fmt.Println(mul)
	fmt.Println(div)
}
