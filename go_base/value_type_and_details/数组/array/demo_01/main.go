// 2023/1/1,17:45
package main

import "fmt"

func main() {
	//arr := []int{5, 233, 324, 470, 599, 697}
	//b := 0
	//for _, v := range arr {
	//	b += v
	//}
	//fmt.Println(b / 6)

	var intArr [4]int
	//int8 占1个字节
	//int16占2个
	//int32占4个
	//int64占8个
	//int   8个
	fmt.Println(intArr)
	//数组第一个值的地址和数组的地址是相同的
	//%p打印地址
	fmt.Printf("    数组地址：%p\n", &intArr)
	fmt.Printf("  数组首地址：%p\n", &intArr[0])
	//自第二个往后依次加上该类型所占的字节
	fmt.Printf("数组第二地址：%p\n", &intArr[1])
	fmt.Printf("数组第三地址：%p\n", &intArr[2])
	fmt.Printf("数组第四地址：%p", &intArr[3])

}
