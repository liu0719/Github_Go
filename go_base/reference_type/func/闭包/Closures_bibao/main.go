// 2022/12/27,17:59
package main

import "fmt"

// AddUpper 累加器
func AddUpper() func(int) int {
	var n int = 10
	return func(x int) int {
		n = n + x
		return n
	}
}

func main() {
	f := AddUpper()
	fmt.Println(f(1))
}
