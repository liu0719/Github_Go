package main

import (
	"fmt"
	"math/rand"
)

type Cat struct {
	Name string
	Age  int
}

type Person struct {
	Name    string
	Age     int
	Address string
}

func main() {
	//定义一个存放任意数据的管道 3个数据
	allChan := make(chan interface{}, 3)
	allChan <- 10
	allChan <- "tom"
	cat := Cat{
		Name: "小花",
		Age:  10,
	}
	allChan <- cat

	//我们希望获得第三个元素，则将前两个推出
	<-allChan
	<-allChan
	//取出后不能直接调用,默认是空接口类型没有字段
	newcat := <-allChan
	fmt.Println(newcat)
	cat1 := newcat.(Cat)
	fmt.Println(cat1.Name)

	personchan := make(chan Person, 10)
	for i := 0; i < 10; i++ {
		person := Person{
			Name:    "tom",
			Age:     rand.Intn(100),
			Address: "北京市",
		}
		personchan <- person
	}
	for i := 0; i < 10; i++ {
		p := <-personchan
		fmt.Printf("第%v个,Name=%v,Age=%v,Adderss=%v\n", i+1, p.Name, p.Age, p.Address)
	}

	// 管道关闭,可以读取，不能写入
}
