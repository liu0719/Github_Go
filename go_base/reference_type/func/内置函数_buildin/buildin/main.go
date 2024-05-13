// 2023/1/1,13:55
package main

import (
	"fmt"
)

func main() {
	//1.len()

	//2.new  分配置类型
	//fmt.Println("\a")
	num := new(int)
	*num = 100
	fmt.Printf("类型：%T\n值：%v\n地址：%v\n指向的值%v\n",
		num, num, &num, *num)
	//make  分配引用类型

}
