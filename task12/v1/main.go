package main

import "fmt"

type Set struct {
	val map[string]string
}

func NewSet() *Set {
	return &Set{
		val: map[string]string{},
	}
}

func (s *Set) Add(str string) {
	_, ok := s.val[str]

	if ok {
		return
	}

	s.val[str] = str
}

func (s *Set) All() []string {
	vals := make([]string, 0, len(s.val))

	for _, v := range s.val {
		vals = append(vals, v)
	}

	return vals
}

func main() {
	arr := []string{"cat", "cat", "dog", "cat", "tree"}
	set := NewSet()

	for _, v := range arr {
		set.Add(v)
	}

	fmt.Println(set.All())
}
