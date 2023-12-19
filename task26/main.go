package main

import (
	"fmt"
	"strings"
)

func UniqueChars(str string) bool {
	lowStr := strings.ToLower(str)
	chars := strings.Split(lowStr, "")
	countChars := make(map[string]int)

	for _, char := range chars {
		_, ok := countChars[char]

		if ok {
			return false
		} else {
			countChars[char] = 1
		}
	}

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
