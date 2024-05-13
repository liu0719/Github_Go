package main

import "fmt"

func 求函数值(n int) int {
	if n == 1 {
		return 3
	} else {
		return 2*求函数值(n-1) + 1
	}
}

func main() {
	var n int
	var res int
	fmt.Println("请输入你想算的值")
	fmt.Scanln(&n)
	res = 求函数值(n)
	fmt.Println("求出来为", res)

}
