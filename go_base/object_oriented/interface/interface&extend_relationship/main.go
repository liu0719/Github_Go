package main

import "fmt"

//接口和继承的关系（图）
//小猴子继承老猴子的特点，但是小猴子不会像鸟（接口）一样飞
//也不会像鱼（接口）一样游，要想像它们一样飞和游的话，就必须
// 要实现他们的所有方法，然后，才能用他们的特性

/*
	接口和继承解决的问题不同
	继承的价值主要在于，解决代码的复用性和可维护性

	而接口则是设计，设计好各种规范（方法），让其自定义类型去实现这些方法
	继承要满足is-a的关系，而接口则是like-a的关系
	接口在一定程度上可以实现代码解耦
*/
//老猴子
type Mokey struct {
	Name string
}

//鸟
type Bird interface {
	flying()
}

//鱼
type Fish interface {
	swimming()
}

//小猴子
type LitterMonkey struct {
	Mokey //继承
}

func (m *Mokey) climbing() {
	fmt.Println(m.Name, "生来会爬树。。。。")
}

//LittleMonkey实现Bird接口
func (m *Mokey) flying() {
	fmt.Println(m.Name, "通过学习，会飞了")
}

//LittleMonkey实现Fish接口
func (m *Mokey) Swimming() {
	fmt.Println(m.Name, "通过学习，会游泳了")
}
func main() {
	monkey := LitterMonkey{
		Mokey{
			Name: "悟空",
		},
	}
	monkey.climbing()
}

/*
1)当A结构体继承了B结构体，那么A结构就自动的继承了B结构体的字段和方法，并且可以直接使用
2)当A结构体需要扩展功能，同时不希望去破坏继承关系，则可以去实现某个接口即可，因此我们可以认为。实现接口是对继承机制的补充.

*/
