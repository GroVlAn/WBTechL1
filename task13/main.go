package main

import "fmt"

func main() {
	a, b := 15, 94

	fmt.Println(a, b)
	// математическое решение
	a += b
	b = a - b
	a = a - b

	fmt.Println(a, b)

	// решение по средствам языка

	a, b = b, a

	fmt.Println(a, b)
}
