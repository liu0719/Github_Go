// 2022/12/28,12:01
package main

import "fmt"

func sum(n1 int, n2 int) {
	//当执行到defer时，暂时不执行，将defer语句压入特殊的栈，defer栈
	//函数执行完毕，按照先入后出的方式，在执行
	//压入栈时，数据的值也会进入栈
	defer fmt.Println(n1)
	defer fmt.Println(n2)
	res := n1 + n2
	fmt.Println(res)
}
func main() {
	sum(2, 3)
}
