package main

import (
	"fmt"
	"runtime"
)

func main() {
	//获取机器当前的逻辑cpu个数
	num := runtime.NumCPU()
	fmt.Println(num)

	//可以自己使用设置多少个cpu
	runtime.GOMAXPROCS(num - 1)
	fmt.Println("ok")
}
