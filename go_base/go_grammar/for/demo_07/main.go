package main

import (
	"fmt"
)

func main() {
	//for i := 0; i < 4; i++ {
	//	for j := 0; j < 10; j++ {
	//		if j == 2 {
	//			continue
	//		}
	//		fmt.Println("j=", j)
	//	}
	//}

	//打印奇数
	//for i := 0; i < 100; i++ {
	//	if i%2 == 0 {
	//		continue
	//	}
	//	fmt.Print(i, " ")
	//}

	//
	//for {
	//	var a int
	//	fmt.Println("请输入一个整数")
	//	fmt.Scanln(&a)
	//	if a > 0 {
	//		fmt.Println("正数")
	//	} else if a < 0 {
	//		fmt.Println("负数")
	//	} else if a == 0 {
	//		fmt.Println("结束")
	//		break
	//	}
	//}
	var positiveCount int = 0
	var negitiveCount int = 0
	var num int
	for {
		fmt.Println("请输入一个整数")
		fmt.Scanln(&num)
		if num > 0 {
			positiveCount++
			continue
		}
		if num < 0 {
			negitiveCount++
			continue
		}
		if num == 0 {
			fmt.Println("结束")
			break
		}
	}
	fmt.Printf("正数是%v个，负数是%d个", positiveCount, negitiveCount)

}
