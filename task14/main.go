package main

import (
	"fmt"
	"reflect"
)

func checkType(val interface{}) {
	valType := reflect.TypeOf(val)

	// для проверки типа в рантайме используем пакет reflect
	switch valType.Kind() {
	case reflect.Int:
		fmt.Printf("%d is integer\n", val)
	case reflect.String:
		fmt.Printf("%s is string\n", val)
	case reflect.Bool:
		fmt.Printf("%v is boolean\n", val)
	case reflect.Chan:
		fmt.Printf("%v is channel\n", val)
	default:
		fmt.Println("unexpected type")
	}
}

func main() {
	var val interface{} = false

	checkType(val)
}
