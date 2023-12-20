package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "snow dog sun"
	var res strings.Builder
	// разбиваем строку на слова
	words := strings.Split(str, " ")

	// собираем новую строку из слов в обратном порядке
	for i := len(words) - 1; i >= 0; i-- {
		res.WriteString(words[i])
		if i != 0 {
			res.WriteString(" ")
		}
	}

	fmt.Println(res.String())
}
