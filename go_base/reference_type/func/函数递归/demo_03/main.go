package main

import "fmt"

func peach(n int) int {
	if n > 10 || n < 1 {
		fmt.Println("输入错误")
		return 0
	}
	if n == 10 {
		return 1
	} else {
		return (peach(n+1) + 1) * 2
	}
}

func add(n *int) {
	*n = *n + 10
}

//函数内部局部变量不影响全局变量
//可以转参数改变也可以传地址改变

func main() {
	a := peach(1)
	fmt.Println(a)
	b := 9
	add(&b)
	fmt.Println(b)
}
