package main

import "fmt"

type Point struct {
	x int
	y int
}

func main() {
	var a interface{}
	point := Point{1, 3}
	a = point
	var b Point
	/*
		这里不能直接赋值需要一个类型断言，
		Point实现了a这个空接口，所以可以转成空接口，而别的结构体也可以实现空接口，
		部分一定属于整体，但整体不一定属于部分
	*/
	// b=a不行
	//b=a.(Point) 类型断言
	b = a.(Point)
	fmt.Println(b)
	var c interface{}
	var d float64 = 1.001
	c = d            //整体包括部分
	y, ok := c.(int) //部分不一定包括整体，所以需要类型断言，判断该值是属于该部分后，才进行转换
	if ok {
		fmt.Println("success")
		fmt.Println(y)
	} else {
		fmt.Println("fail")
	}

}
