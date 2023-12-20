package main

import (
	"fmt"
	"strings"
)

func UniqueChars(str string) bool {
	// приводим строку к нижнему регистру и разбиваем её на символы
	lowStr := strings.ToLower(str)
	chars := strings.Split(lowStr, "")
	countChars := make(map[string]int)

	// записываем в map сиволы и количество их в искомой строке, если же такой символ уже есть в map
	// возвращаем false
	for _, char := range chars {
		_, ok := countChars[char]

		if ok {
			return false
		} else {
			countChars[char] = 1
		}
	}

	// если в map все элементы встречаются 1 раз, возвращаем true
	return true
}

func main() {
	str1 := "abcd"
	str2 := "abCdefAaf"
	str3 := "aabcd"

	fmt.Printf("string 1: %v \n", UniqueChars(str1))
	fmt.Printf("string 2: %v \n", UniqueChars(str2))
	fmt.Printf("string 3: %v \n", UniqueChars(str3))
}
