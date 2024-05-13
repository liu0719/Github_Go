// 2023/1/2,18:59
package main

import (
	"fmt"
)

func main() {
	//1.string底层是一个byte数组，也可以进行切片梳理
	//2.string内存图
	//3.string本身不可变
	str := "hello world"
	sli := str[6:]
	fmt.Println(sli)

	//4.要修改字符串要先转成byte[]或rune[]进行修改
	//修改后再转成string
	str1 := []rune(str)
	str1[2] = '安'
	str = string(str1)
	fmt.Println(str)

	//细节：转成byte后可以处理英文和数字
	//byte是按字节处理汉字一个字站三个字节，会出现乱码
	//解决方法是转成rune[]rune[]是int32的别名
}
