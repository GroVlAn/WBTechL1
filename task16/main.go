package main

import "fmt"

func quickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	res := make([]int, 0)
	// находим индекс опорного элемента, в данном случай центральный
	mid := arr[len(arr)/2]
	// два слайса, для хранения левой и правой части слайса
	var less, greater []int

	// заполняем правые и левые части массива
	// в левую те что, меньше опорного, в правую, те что больше опорного
	for i := 0; i < len(arr); i++ {
		if i == len(arr)/2 {
			continue
		}

		if arr[i] < mid {
			less = append(less, arr[i])
		} else {
			greater = append(greater, arr[i])
		}

	}

	// рекурсивно вызываем сортировку для левой и правой части и объединяем в результирующий слайс
	res = append(quickSort(less), mid)
	res = append(res, quickSort(greater)...)

	return res
}

func main() {
	arr := []int{10000, -1, 0, -1223123, -1, 1222222, 134}

	fmt.Println(quickSort(arr))
}
