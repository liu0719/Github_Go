// 2023/1/3,20:21
package main

import (
	"fmt"
	"math"
)

func main() {
	//1
	//fmt.Println("1---------------------------")
	//arr := [10]int{}
	//rand.Seed(time.Now().UnixNano())
	//for i := 0; i < len(arr); i++ {
	//
	//	arr[i] = rand.Intn(101)
	//}
	//max, min, flag := 0, arr[0], false
	//for i := len(arr) - 1; i > 0; i-- {
	//	fmt.Println(arr[i])
	//	if arr[i] > max {
	//		max = arr[i]
	//	}
	//	if arr[i] < min {
	//		min = arr[i]
	//	}
	//	if arr[i] == 55 {
	//		flag = true
	//	}
	//}
	//fmt.Println(max, min, flag)
	fmt.Println("2----------------------")
	//arr2 := []int{1, 2, 3, 4, 5, 7, 8, 9}
	//list := []int{}
	//num := 6
	//for i := range arr2 {
	//	if arr[i] > num {
	//		for j := 0; j < i-1; j++ {
	//			list[j] = arr[i]
	//		}
	//		list[i] = num
	//		for j := i + 1; j < len(arr); j++ {
	//			list[j] = arr[j-1]
	//		}
	//	}
	//}
	//fmt.Println(list)
	//10
	num := 0
	count := [8]int{}
	for i := range count {
		fmt.Println("请输入第", i+1, "个评委的分数")
		_, _ = fmt.Scanln(&count[i])
		num += count[i]
	}
	fmt.Println("输入完成，正在计算。。。")
	result := num / len(count)
	for i, v := range count {
		count[i] = int(math.Abs(float64(result - v)))
	}
	good, bad := 0, 0
	for i, v := range count {
		if v > count[bad] {
			bad = i
		}
		if v < count[good] {
			good = i
		}
	}
	fmt.Printf("计算结果为:\n 该选手分数为：%d\n最佳评委是:%d\n最差评委是：%d\n",
		result, good+1, bad+1)
}
