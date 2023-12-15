package main

import "fmt"

type Human struct {
	name string
	age  int
}

func (h *Human) GrowUp() {
	if h.age < 100 {
		h.age++
	}
}

func (h *Human) ChangeName(newName string) {
	h.name = newName
}

type Action struct {
	Human
}

func main() {
	action := new(Action)

	action.name = "Ivan"
	action.age = 17

	fmt.Printf("name: %s \n age: %d \n", action.name, action.age)

	action.GrowUp()
	action.ChangeName("Sergey")

	fmt.Printf("name: %s \n age: %d \n", action.name, action.age)
}
