// 2023/1/1,18:26
package main

import (
	"fmt"
	"time"
)

func main() {
	var score [5]float64
	fmt.Println("请分别输入五人的成绩")
	time.Sleep(time.Second * 2)
	for i := 0; i < len(score); i++ {
		fmt.Println("请输入第", i+1, "个人的成绩")
		fmt.Scanln(&score[i])
	}
	fmt.Printf("五人的成绩分别为%v\n", score)

	//    四种初始化数组方式
	var array01 [2]int
	fmt.Println(array01)
	var array02 = [...]int{2, 3, 4}
	fmt.Println(array02)
	var array03 = []int{3, 4, 5} //三个点可以省略不写
	fmt.Println(array03)
	var array04 [3]int = [3]int{2, 3, 4}
	fmt.Println(array04)
	//指定元素下标
	var array05 = []int{3: 0, 2: 1, 1: 9, 0: 100}
	fmt.Println(array05)
}
