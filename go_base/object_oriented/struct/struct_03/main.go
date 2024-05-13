// 2023/1/5,9:49
package main

import (
	"encoding/json"
	"fmt"
)

//字段完全相同的结构体，可以相互转换，否则不能转换
type Stu struct {
	Name int
}

type Person struct {
	Name int
}

type A struct {
	num int
}
type B A

type Monsters struct {
	Name  string `json:"name"` //`json:"name"`就是struct tag
	Age   int    `json:"age"`
	Skill string `json:"skill"`
}

func main() {
	a := Stu{2}
	b := Person{2}
	fmt.Println(a == Stu(b))
	b2 := B{3}
	fmt.Println(b2)
	a2 := A{2}
	//不能直接进行赋值，是不一样的结构体
	//需要强制转换后才能赋值
	b2 = B(a2)
	fmt.Println(b2)
	monster := Monsters{"牛魔王", 500, "nmnmn"}
	fmt.Println(monster)
	//json.Marshal()用到了反射
	jsonStr, err := json.Marshal(monster)
	if err != nil {
		fmt.Println("错误")
	}
	fmt.Println(string(jsonStr))
}
