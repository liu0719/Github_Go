package main

import (
	"fmt"
)

//接口注意事项
/*
1.接口本身不能创建实例,但是可以指向一个实现了该接口的自定义类型的变量（实例）
2.接口中所有的方法都没有方法体,即都是没有实现的方法
3.golang中一个自定义类型需要将某个接口的的所有方法都实现，才能说这个自定义类型实现了该接口
4.一个自定义类型只有实现了某个接口，才能将盖子定义变量的实例（变量）赋给接口类型
5.只要是自定义数据类型，就可以实现接口，不仅仅只是结构体
6.一个自定义类型可以实现多个接口
7.接口中不能有任何自定义变量
8一个接口(A)可以集成多个别的接口（B C），这时如果呀哦实现接口A也必须将A继承的BC接口的方法也全部实现
9.interface 默认是一个指针类型（引用类型）如果没有对interface进行初始化就使用就会输出nil
10.空接口没有任何方法所以所有类型都实现了空接口

*/
// 4 eg
type Stu struct {
}

func (s Stu) haha() {
	fmt.Println("haha")
}

// 5 eg
//可以理解为给int起了一个别名，也是自定义数据类型
type integer int

func (i integer) haha() {
	fmt.Println("haha", i)
}

type AI interface {
	haha()
}

//
type A interface {
	text01()
	text02()
}
type B interface {
	test01()
	test03()
}
type C interface {
	A
	B
}

func main() {

	var stu Stu //结构体变量，实现了say()实现了AI
	var a AI = stu
	a.haha()
	var i integer = 10
	var b AI = i
	b.haha()
}
