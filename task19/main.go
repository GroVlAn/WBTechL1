package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "главрыба"
	var res strings.Builder
	chars := strings.Split(str, "")

	for i := len(chars) - 1; i >= 0; i-- {
		res.WriteString(chars[i])
	}

	fmt.Println(res.String())
}
