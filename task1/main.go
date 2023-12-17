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

// Action встраивание полей и методов структуры Human в структуру Action
type Action struct {
	Human
}

func main() {
	// новый экземпляр структуры Action
	action := new(Action)

	// так как структура Human была встроена в Action, мы имеем те же поля, что и у Human
	action.name = "Ivan"
	action.age = 17

	fmt.Printf("name: %s \n age: %d \n", action.name, action.age)

	// а так же мы имеем методы структуры Human
	action.GrowUp()
	action.ChangeName("Sergey")

	fmt.Printf("name: %s \n age: %d \n", action.name, action.age)
}
