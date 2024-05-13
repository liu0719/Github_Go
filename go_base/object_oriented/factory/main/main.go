// 2023/1/7,13:01
package main

import (
	"fmt"
	"go_base/object_oriented/factory/model"
)

func main() {
	s1 := model.Student{Name: "xiaoming", Score: 100.0}
	fmt.Println(s1)
	//结构体首字母小写，可以通过工厂模式来解决
	s2 := model.NewStudent("小明", 99.9)
	fmt.Println(*s2)
	fmt.Println("name=", s2.Name, "score=", s2.GetScore())
}
