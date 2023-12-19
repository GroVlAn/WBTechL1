package main

import "strings"

// в данном примере используется strings.Builder вместо string, так как он будет рассширятся только при добавлении
// данных. При использовании string вся строка хранится в памяти, хоть мы и используем только 100 символов
var justString strings.Builder

func someFunc() {
	v := createHugeString(1 << 10)
	justString.WriteString(v[:100])
}

func main() {
	someFunc()
}
