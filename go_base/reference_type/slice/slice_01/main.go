// 2023/1/1,22:25
package main

import (
	"fmt"
)

func main() {
	//切片基本使用
	arr := [39]int{1, 2, 3, 4, 5, 6, 5, 4, 3, 3, 2, 2}
	slice := arr[0:11]      //包括1不包括4
	fmt.Println(slice)      //切片元素
	fmt.Println(len(slice)) //切片长度
	//一般等于切的数组的长度
	fmt.Println(cap(slice)) //切片容量是可以动态变化的
	//切片是引用类型，传递的是指针地址，
	//修改数值后会引起底层数据发生变化

	//切片使用的三种方式
	//1.创建数组用切片去引用创建好的数组
	//2.make创建
	//切片名称:= make(类型，长度，容量（可选）)
	fmt.Println("2---------------------")
	slice1 := make([]int, 4, 10)
	fmt.Println(slice1) //默认值都为 0
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))
	//3.定义切片直接赋值具体数组，原理与make相似
	fmt.Println("3-----------------------")
	slice2 := []int{23, 4, 5, 56, 6}
	fmt.Println("元素", slice, "长度", len(slice2), "容量", cap(slice2))
	slice3 := slice2[:3] //0-3不包括3
	fmt.Println(slice3)
	slice4 := arr[2:] //切数组的话是2到末尾  切切片的话是只取2下表的元素
	fmt.Println(slice4)
	slice5 := slice2[:] //全取 [0:len(slice2)]
	fmt.Println(slice5)

	//！！！！！！！！！！！！！！！
	//引用类型，值变全变！！！！！！！！！！！
	//！！！！！！！！！！！！！！！！！！！！！！！！

	//append内置函数
	fmt.Println("append---------------------")
	sli := []int{100, 200, 300}
	sli = append(sli, 400)
	fmt.Println(sli)
	sli = append(sli, slice2...)
	fmt.Println(sli)
	slice6 := []int{1, 2, 3, 4, 5, 6}
	var slice7 = make([]int, 10, 10)
	copy(slice7, slice6) //必须都为切片才能拷贝
	fmt.Println(slice7)
}
