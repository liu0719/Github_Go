package main

import "fmt"

//全局匿名函数
var (
	Fun1 = func(a int) int {
		return a + 10
	}
)

func main() {
	a := func(n1 int, n2 int) int {
		return n1 + n2
	}(10, 20)
	fmt.Println(a)
	b := func(n1 int, n2 int) int {
		return n1 - n2
	}
	fmt.Println(b(1, 2))
	//	全局的使用
	fmt.Println(Fun1(2))
}
