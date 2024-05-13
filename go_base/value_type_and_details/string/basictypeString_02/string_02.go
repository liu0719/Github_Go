package main

import (
	"fmt"
	"strconv"
)

func main() {
	var str string = "true"
	var b bool
	//strconv包
	//ParseBool 返回两个值_忽略一个错误
	b, _ = strconv.ParseBool(str)
	//b = strconv.ParseBool(str1)
	fmt.Printf("b string:%T b=%v", b, b)
	var str1 string = "4565890"
	var n1 int64
	var n2 int
	fmt.Println()
	n1, _ = strconv.ParseInt(str1, 10, 64)
	n2 = int(n1)
	fmt.Printf("n1 string=:%T,n1=%v", n2, n2)
	fmt.Println()

	fmt.Printf("n2 string:%T,n2=%v", n2, n2)
	fmt.Println()
	var str3 string = "123.456"
	var f1 float64
	//go语言不能完成转换时，变量重新赋值为默认值
	// “hello”转 int型
	f1, _ = strconv.ParseFloat(str3, 64)
	fmt.Printf("f1 string:%T,f1=%v", f1, f1)
}
