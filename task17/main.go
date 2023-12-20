package main

import "fmt"

func quickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	res := make([]int, 0)
	mid := arr[len(arr)/2]
	var less, greater []int

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

	res = append(quickSort(less), mid)
	res = append(res, quickSort(greater)...)

	return res
}

func binarySearch(arr []int, n int) int {
	// инициализируем правую и левую границу слайса
	left := 0
	right := len(arr) - 1

	// пока левая граница меньше правой ищем элемент
	for left <= right {
		// индекс центрального элемента слайса
		mid := (left + right) / 2

		// если нашли, возвращаем индекс элемента
		if arr[mid] == n {
			return mid
		}

		// если центральный элемент меньше искомого сдвигаем левую часть, иначе сдвигаем правую
		if arr[mid] < n {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	// если не нашли, возвращаем -1
	return -1
}

func main() {
	arr := []int{10000, -1, 0, -1223123, -1, 1222222, 134, 1}

	arr = quickSort(arr)
	fmt.Println(binarySearch(arr, 134))
}
