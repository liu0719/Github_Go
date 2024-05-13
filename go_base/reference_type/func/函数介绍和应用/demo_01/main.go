package main

import "fmt"

func main() {
	var n1 float64
	var n2 float64
	var operator byte = '+'
	fmt.Println("请输入第一个值")
	fmt.Scanln(&n1)
	fmt.Println("请输入第二个值")
	fmt.Scanln(&n2)
	a := Cal(n1, n2, operator)
	fmt.Println(a)

}
func Cal(n1 float64, n2 float64, operator byte) float64 {
	var res float64
	switch operator {
	case '+':
		res = n1 + n2
	case '-':
		res = n1 - n2
	case '*':
		res = n1 * n2
	case '/':
		res = n1 / n2
	default:
		fmt.Println("操作符错误")
	}
	return res
	// utils.go专门定义函数，让其他文件调用的工具
	// db.go用于定义数据库操作的函数
	//	同一个包里不能有相同的函数和变量名
}
