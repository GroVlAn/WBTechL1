package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "snow dog sun"
	var res strings.Builder
	words := strings.Split(str, " ")

	for i := len(words) - 1; i >= 0; i-- {
		res.WriteString(words[i])
		if i != 0 {
			res.WriteString(" ")
		}
	}

	fmt.Println(res.String())
}
