// 2022/12/30,12:09
package main

import (
	"fmt"
)

func printPyramid() {
	var n int
	fmt.Println("请输入打印的行数")
	fmt.Scanf("%d", &n)
	for i := 0; i < n; i++ {
		for j := 0; j < n-1-i; j++ {
			fmt.Print(" ")
		}
		for k := 0; k < 2*i+1; k++ {
			if k == 0 || k == 2*i {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}

		}
		fmt.Println()
	}

	for i := 0; i < n-1; i++ {
		for j := 0; j < i+1; j++ {
			fmt.Print(" ")
		}
		for k := 0; k < (n-1-i)*2-1; k++ {
			if k == 0 || k == (n-1-i)*2-2 {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
func main() {
	printPyramid()
}
