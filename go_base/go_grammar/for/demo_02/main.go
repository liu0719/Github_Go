package main

import "fmt"

func main() {
	// num := 0
	// for i := 0; i < 100; i++ {
	// 	if i%9 == 0 {
	// 		fmt.Println(i)
	// 		num += i
	// 	}
	// }
	// fmt.Println("综合：", num)

	// 2.
	var n int = 64
	for i := 0; i < n+1; i++ {
		fmt.Println(i, " + ", n-i, " = ", n)
	}
	// go没有while和dowhile
	// 可以用for+break实现
	// while
	// for{
	// 	if  {
	// 		break
	// 	}
	// 	// 要执行的代码
	// }

	// dowhile
	// for{
	// 	// 要执行的代码
	// 	 if {
	// 		break
	// 	}
	// }

}
