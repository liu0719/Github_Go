// 2022/12/27,18:18
package main

import (
	"fmt"
	"strings"
)

func makeSuffix(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}
func main() {

	f := makeSuffix(".json")
	fmt.Println(f("name"))
	fmt.Println(f("hello.json"))
	fmt.Println(f("index.html"))
	//	闭包只需要传入一次后缀
}
