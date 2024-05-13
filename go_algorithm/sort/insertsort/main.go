package main

import "fmt"

//插入排序
func InsertSort(arr *[]int) {
	//完成第一次

	for i := 1; i <= len(*arr)-1; i++ {
		insertVal := (*arr)[i]
		insertIndex := i - 1
		for insertVal >= 0 && (*arr)[insertIndex] < insertVal {

			(*arr)[insertIndex+1] = (*arr)[insertIndex] //数据向后移一位
			insertIndex--
			if insertIndex == -1 {
				(*arr)[insertIndex+1] = insertVal
				break
			}
		}
		//插入
		if insertIndex+1 != i {
			(*arr)[insertIndex+1] = insertVal
		}
		fmt.Println("第", i, "次插入后:", *arr)
	}
}
func main() {
	arr := []int{23, 0, 12, 56, 34, 100, 999, 444, 556, 779, 867, 245, 2, 40, 23}
	InsertSort(&arr)
}
