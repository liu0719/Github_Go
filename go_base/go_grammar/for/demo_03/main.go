package main

import (
	"fmt"
	_ "go/printer"
)

func main() {
	// for j := 0; j < 3; j++ {
	// 	num := 0.0
	// for i := 0; i < 5; i++ {
	// 	var a float64
	// 	fmt.Printf("请输入第%d个班级第%d个学生的成绩\n", j+1,i+1)
	// 	fmt.Scanln(&a)
	// 	num += a
	// }
	// fmt.Printf("第%d个班级的平均分%f \n", j+1,num/5)
	// }
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
	// for i := 0; i < n; i++ {
	// 	for j := 0; j < i+1; j++ {
	// 		fmt.Print(" * ")
	// 	}
	// 	fmt.Println()
	// }

}
