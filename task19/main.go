package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "главрыба"
	var res strings.Builder
	// разбиваем строку на символы
	chars := strings.Split(str, "")

	// собираем символы в обратном порядке в новую строку
	for i := len(chars) - 1; i >= 0; i-- {
		res.WriteString(chars[i])
	}

	fmt.Println(res.String())
}
