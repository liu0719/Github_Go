package main

import "fmt"

func runningSum(nums []int) []int {
	for i := range nums {
		if i != 0 {
			nums[i] = nums[i-1] + nums[i]
		}
	}
	return nums
}

func main() {
	// for i := 0; i < 10; i++ {
	// 	fmt.Println("你好，尚硅谷")
	// }

	// 第二种写法
	// var j=1
	// for j<=10{
	// 	j++
	// }

	// 死循环  配合break使用
	// for{

	// }

	// 出现中文的话会乱码，中文三个字节，而他是一个字节遍历的
	var str string = "hello world!北京"
	for index, val := range str {
		fmt.Printf("inedx=%d,val=%c", index, val)
		fmt.Println()
	}

}
