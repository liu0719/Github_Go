package main

import "fmt"

func main() {
	a := 0
	b := 0
	fmt.Println("请输入第一个数：")
	fmt.Scanln(&a)
	fmt.Println("请输入第二个数：")
	fmt.Scanln(&b)
	c := 我是函数(a, b)
	fmt.Println("结果是：", c)
}
func 我是函数(a, b int) int {
	return a + b
}
