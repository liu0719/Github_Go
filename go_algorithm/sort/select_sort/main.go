package main

import "fmt"

func SelectSort(arr *[]int) {
	//标准的访问方式
	//标准(*arr)[1] = 600
	// (*arr)[1] = 200
	// 假设arr[0]为最大值
	max := (*arr)[0]
	maxIndex := 0
	// // 遍历后面1----len(arr)-1
	for i := 0; i < len(*arr)-1; i++ {
		max = (*arr)[i]
		maxIndex = i
		// 遍历后面1----len(arr)-1
		for j := i + 1; j < len(*arr); j++ {
			if max < (*arr)[j] { //找到真正的最大值
				max = (*arr)[j]
				maxIndex = j
			}
		}
		if maxIndex != 1 {
			(*arr)[i], (*arr)[maxIndex] = (*arr)[maxIndex], (*arr)[i]
		}
		fmt.Println("第", i+1, "次，", arr)
	}
}

func main() {
	//定义数组
	arr := []int{10, 34, 19, 100, 80, 130, 220}
	SelectSort(&arr)
	fmt.Println(arr)
}
