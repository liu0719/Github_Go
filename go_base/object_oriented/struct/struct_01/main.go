// 2023/1/4,15:42
package main

import "fmt"

//结构体是自定义的数据类型，代表一类事物
type Cat struct {
	Name  string
	Age   int
	Color string
}

//结构体是值类型！！！！！！！！！！！！！！
func main() {
	var cat1 Cat
	cat1 = Cat{"小花", 10, "yellow"}
	fmt.Println(cat1)
	fmt.Printf("%p\n", &cat1)
}
