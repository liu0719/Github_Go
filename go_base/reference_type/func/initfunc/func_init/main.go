package main

//引入包

import (
	"fmt"
	"go_base/reference_type/func/initfunc/utils"
)

func text() {
	fmt.Println("text")
}

func init() {
	text()
	fmt.Println("init")

}

func main() {
	fmt.Println("mian")
	//先执行引入的包里的init函数
	fmt.Println(utils.Age)
}

//init初始化函数，在main前执行
//如果init调用函数text函数,那么先执行text函数
