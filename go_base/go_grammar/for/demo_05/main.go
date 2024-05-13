package main

import "fmt"

func main() {
	//	跳出最外层循环，内部循环不再执行
	//	break可以指定标签（标签自己定义）跳出对应标签的for循环
label1:
	for i := 0; i < 4; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				break label1
			}
			fmt.Println("j=", j)
		}
	}
}
