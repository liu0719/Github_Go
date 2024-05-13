package main

import "fmt"

func addUpper(n int) (a int) {
	for i := 1; i <= n; i++ {
		a += i
	}
	return
}

func main() {
	//传统的测试方式，就是在main函数中调用看看结果是否正确
	s := addUpper(10)
	fmt.Println(s)
}
