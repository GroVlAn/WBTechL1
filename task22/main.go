package main

import (
	"fmt"
	"math/big"
)

func main() {
	// объявляем переменные для хранения суммы, разницы, частного и делимого типа big.Int
	sum := new(big.Int)
	sub := new(big.Int)
	mul := new(big.Int)
	div := new(big.Int)
	// два значения типа big.Int, которые будем получать из консоли
	num1 := new(big.Int)
	num2 := new(big.Int)

	fmt.Print("Enter number 1: ")
	fmt.Scan(num1)

	fmt.Print("Enter number 2: ")
	fmt.Scan(num2)

	// Выполняем сложение, вычитание, умножение используя big.Int
	sum.Add(num1, num2)
	sub.Sub(num1, num2)
	mul.Mul(num1, num2)

	// проверяем на ноль, делитель, если не раве, то выполняем деление, иначе выводим ошибку
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
