// 2023/1/1,14:50
package main

import (
	"errors"
	"fmt"
	"time"
)

func text() {
	//错误处理
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
			fmt.Println("发送邮件给管理员")
		}
	}()
	//
	num := 0
	num1 := 10
	res := num1 / num
	fmt.Println(res)
}

//函数去读取配置文件的信息
func readConf(name string) (err error) {
	if name != "index.html" {
		//读取
		return nil
	} else {
		//返回自定义错误
		return errors.New("读取文件错误")
	}
}

//自定义错误
func text02() {
	err := readConf("index.htm")
	if err != nil {
		//如果读取文件发送错误，就输出这个错误，并终止程序
		panic(err)

	}
	fmt.Println("text02继续执行")
}
func main() {
	//text()
	//fmt.Println("错误机制")
	text02()
	fmt.Println("main...")
	time.Sleep(8 * time.Second)
}
