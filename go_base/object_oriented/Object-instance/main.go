// 2023/1/6,19:32
package main

import (
	"fmt"
)

// Student 1)声明(定义)结构体，确定结构体名
//2)编写结构体的字段
//3)编写结构体的方法10.3.2学生案例;
//1)编写一个Student结构体，包含name、 gender、age、id、score字段，分别为 string 、string、int、int、 float64类型。
//2)结构体中声明一个say方法，返回string 类型，方法返回信息中包含所有字段值。
//3)在main方法中，创建Student结构体实例(变量)，并访间 say方法，并将调用结果打印输出。
//4)走代码
type Student struct {
	Name   string
	Gender string
	Age    int
	Id     int
	Score  float64
}

func (s Student) say() string {
	return fmt.Sprintf("我的名字是%v,性别是%v,%v岁了,学号ID是：%v,成绩：%v",
		s.Name, s.Gender, s.Age, s.Id, s.Score)
}

type Box struct {
	length float64
	width  float64
	height float64
}

func (b Box) BoxV() float64 {
	return b.width * b.length * b.height
}

type Visitor struct {
	Name string
	Age  int
}

func (v *Visitor) showPrice() {
	if v.Age > 80 || v.Age < 8 {
		fmt.Printf("%v,滚蛋不让玩\n", v.Name)
	} else if v.Age > 18 {
		fmt.Printf("%v,成人价 20元\n", v.Name)
	} else {
		fmt.Printf("%v,儿童价 免费\n", v.Name)
	}

}
func main() {
	s1 := Student{"小明", "男", 20, 2021030441012, 100.99999}
	fmt.Println(s1.say())
	box1 := Box{10.32, 3.43, 1.89}
	fmt.Println(box1.BoxV())
	var v Visitor
	fmt.Println(v)
	//for {
	//	fmt.Println("请输入你的名字")
	//	fmt.Scanln(&v.Name)
	//	if v.Name == "n" {
	//		fmt.Println("exit...")
	//		break
	//	}
	//	fmt.Println("请输入你的年龄")
	//	fmt.Scanln(&v.Age)
	//	v.showPrice()
	//	time.Sleep(time.Second * 2)
	//}
	//这种声明方式不受字段顺序影响
	v1 := Visitor{
		Name: "ahaa",
		Age:  100,
	}
	fmt.Println(v1)
	v2 := &Visitor{"小明", 23}
	v3 := &Visitor{
		Age:  30,
		Name: "校历",
	}
	fmt.Println(v2, v3)   //&{小明 23} &{校历 30}
	fmt.Println(*v2, *v3) //{小明 23} {校历 30}
	//*指针的取值运算符
}
