package main

import "fmt"

//快速排序
func QuickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	pivot := arr[0]
	var left, right []int

	for _, ele := range arr[1:] {
		if ele <= pivot {
			left = append(left, ele)
		} else {
			right = append(right, ele)
		}
	}
	return append(QuickSort(left),
		append([]int{pivot}, QuickSort(right)...)...)
}

func main() {
	arr := []int{-9, 78, 0, 23, -567, 70}
	arr = QuickSort(arr)
	fmt.Println(arr)
}
