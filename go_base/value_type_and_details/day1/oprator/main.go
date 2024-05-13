package main

import (
	"fmt"
	"runtime"
	"strconv"
)

func 我是函数() {
	var a float32 = 189
	var b float32 = 3
	var c float32 = a + b
	var d float32 = a - b
	var e float32 = a * b
	var f float32 = a / b

	var m int = 8
	var n int = 3
	var g = m % n

	fmt.Printf("加法：c=%f\n", c) //11
	fmt.Println("减法：d:", d)    //5
	fmt.Println("乘法：e:", e)    //24
	fmt.Println("除法：f:", f)    //3.6
	fmt.Println("取余：g:", g)    //2
	if a > 10 {
		fmt.Println("a>10")
	}

}
func bit_op() {
	fmt.Printf("os arch %s,int size %d\n", runtime.GOARCH, strconv.IntSize)
	fmt.Printf("260 %b\n", 260) //2的8次方+2的2次方
	fmt.Printf("%b\n", 6|5)     //二进制取&  110 & 101 =100
	// 二进制取| 110 | 101 = 111
	//二进制取^  110 ^ 101= 011
	// 5<<1 5的二进制向左移一位 101<<1010
	// b++单独一行不支持++b
	b := 7
	a := 5
	b ^= a
	fmt.Println(b)
}
func zhizhen() {
	a := 5
	ptr := &a
	fmt.Printf("%p %d %p", ptr, ptr, &a)
}

var c = 8 //全局变量不能用：=
// 小写字母不能挎包引用
func main() {
	// 我是函数()
	// bit_op()
	// b := util.BinaryFormat(260)
	// c := util.BinaryFormat(-260)
	// fmt.Println(c)
	// fmt.Println(b)
	// fmt.Println("=========================")
	// zhizhen()
	arr := []int{2, 7, 11, 15}
	fmt.Println(arr)

}
