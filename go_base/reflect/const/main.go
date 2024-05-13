package main

import (
	"fmt"
)

const (
	a = iota //从零开始，每行自增1
	b
	c
	d
	e
	f
)

func main() {
	var num int = 0
	num = 90
	//常量使用const
	// 常量在定义时必须初始化,必须付给确定的值，也不能是变量

	const tax int = 10
	//常量不能修改
	//tax=100  cannot assign to tax (constant 10 of type int)
	fmt.Println(num, tax)
	fmt.Println(a, b, c, d, e, f)
}
