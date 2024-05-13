// 2023/1/4,18:06
package main

import "fmt"

//结构体变量（实例）代表一个具体的变量
//结构体内存布局(图片)
//如果结构体的字段类型为指针，slice,mapm即无法自定分配空间的引用类型
//如果要使用，先make一下才能使用

type Person struct {
	Name   string
	Age    int
	Scores [5]float64
	ptr    **int             //指针
	slice  []int             //切片
	map1   map[string]string //map
}

//不同结构体间，字段相互独立，互不影响
//结构体是值类型！！！！！！！！！！！！！！！！
//值类型值拷贝，引用类型指针拷贝
type Stu struct {
	Name int
}

func main() {
	a := 100
	b := &a
	p1 := Person{}
	fmt.Println(p1) //没有分配空间就都是nil
	//引用类型要make后才能使用
	//slice要make
	p1.slice = make([]int, 1)
	p1.slice[0] = 99
	fmt.Println(p1)
	//map要make
	p1.map1 = make(map[string]string)
	p1.map1["key1"] = "haha"
	fmt.Println(p1)
	p1.ptr = &b
	fmt.Println(p1)
	//创建结构体方式
	//1直接声明
	var person1 Person
	fmt.Println(person1)
	//2{}
	var stu2 = Stu{10}
	fmt.Println(stu2)
	//3-&
	var stu3 = new(Stu)
	fmt.Printf("%p\n", stu3)
	//等价于p3.Name=100,h会进行处理，加上取值运算
	(*stu3).Name = 100
	fmt.Println(*stu3)

	//4-{}
	//下面也可以直接给字段赋初值
	var stu4 *Stu = &Stu{}
	//因为stu4是一个指针，所以标准访问字段方法
	(*stu4).Name = 88
	//为了使用方便也可以stu4.name=88
	//原因和第三种一样
	fmt.Println(*stu4)

	/*
			fmt.Println(*p2.Age)
			这种方式试错的，.的优先级高于*
			要加上小括号
		fmt.Println((*p2).Age)
	*/
}
