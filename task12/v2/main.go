package main

import "fmt"

type Set struct {
	// в качестве основы берём слайс, но решение не эффективное, так как при проверке наличия значения в множестве,
	// нужно обойти все значения, а значит скорость добавления будет O(N), но без использования map выполнить более
	// эффективное решение невозможно
	values []string
}

func NewSet() *Set {
	return &Set{
		values: make([]string, 0),
	}
}

func (s *Set) Exist(str string) bool {
	for _, exist := range s.values {
		if exist == str {
			return true
		}
	}

	return false
}

func (s *Set) Add(str string) {
	if s.Exist(str) {
		return
	}

	s.values = append(s.values, str)
}

func (s *Set) All() []string {
	return s.values
}

func main() {
	arr := []string{"cat", "cat", "dog", "cat", "tree"}
	set := NewSet()

	for _, v := range arr {
		set.Add(v)
	}

	fmt.Println(set.All())
}
