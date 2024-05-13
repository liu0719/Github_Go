package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num1 int = 99
	var num2 float64 = 23.996
	var b bool = true
	var mychar byte = 'h'
	var str string //空的字符串

	//1.方法一   fmt.Sprintf()转换字符串
	str = fmt.Sprintf("%d", num1)
	fmt.Printf("str string ：%T str：=%q", str, str)
	fmt.Println()

	str = fmt.Sprintf("%.3f", num2)
	fmt.Printf("str string ：%T str=%q", str, str)

	fmt.Println()
	str = fmt.Sprintf("%t", b)
	fmt.Printf("str string : %T str=%q", str, str)

	fmt.Println()
	str = fmt.Sprintf("%c", mychar)
	fmt.Printf("str string:%T str=%q", str, str)
	//	方法二strconv包的函数 strconv.FormatFloat(num1, 1, 1, 1)
	var num3 int = 99
	var num4 float64 = 23.648
	//var b2 bool = true
	fmt.Println()
	str = strconv.FormatInt(int64(num3), 10)
	fmt.Printf("str string=%T str=%v", str, str)

	fmt.Println() //          要转的数    格式     保留几位     float64的类型
	str = strconv.FormatFloat(num4, 'f', 10, 64)
	fmt.Printf("str string:%T str=%q", str, str)

	//itoa()把int类型转成string
	fmt.Println()
	var num5 int = 345
	str = strconv.Itoa(num5)
	fmt.Printf("str string :%T str=%q", str, str)
}
