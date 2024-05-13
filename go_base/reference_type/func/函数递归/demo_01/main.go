package main

import "fmt"

func 斐波那契数(n int) int {
	if n == 1 || n == 2 {
		return 1
	} else {
		return 斐波那契数(n-1) + 斐波那契数(n-2)
	}
}
func main() {
	var res int
	res = 斐波那契数(60)
	fmt.Println("res=", res)
}
