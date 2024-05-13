// 2023/1/3,18:35
package main

import (
	"fmt"
	"time"
)

func main() {
	arr := [4][6]int{}
	arr2 := [...][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(arr)
	fmt.Println(arr2)
	list := [3][5]float64{}
	for i, v1 := range list {
		for j := range v1 {
			fmt.Println("请输入第", i+1, "个班第", j+1, "个学生的成绩")
			fmt.Scanln(&list[i][j])
		}
	}
	fmt.Println("输入完成,正在计算总分和平均分。。。")
	time.Sleep(time.Second * 2)
	for i, v1 := range list {
		count := 0.0
		for _, v2 := range v1 {
			count += v2
		}
		fmt.Println(i+1, "班")
		fmt.Println("总分:", count)
		fmt.Println("平均分:", count/float64(len(list[i])))
	}
}
