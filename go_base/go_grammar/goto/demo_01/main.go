package main

import "fmt"

func main() {
	var i int

la:
	fmt.Println("请输入一个数")
	fmt.Scanln(&i)
	if i != 0 {
		goto la
	}
	// return 在方法或函数表示终止函数
	// main函数中表示终止

}
