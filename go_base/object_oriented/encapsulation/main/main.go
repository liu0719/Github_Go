package main

import (
	"fmt"
	"go_base/object_oriented/encapsulation/model"
)

func main() {
	//person
	p := model.NewPerson("小明")
	p.SetAge(40)
	fmt.Println(p.GetAge())
	p.SetSal(8888.88)
	fmt.Println(p.GetSal())
	p.SetSal(9999.99)
	fmt.Println(p.GetSal())

	// account
	a := model.NewAccount("lgs1112", "liu019", 30)
	if a != nil {
		fmt.Println("创建成功", *a)
	} else {
		fmt.Println("创建失败", a)
	}
}
