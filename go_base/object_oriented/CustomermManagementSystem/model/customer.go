// 负责数据
package model

import "fmt"

//声明一个结构体，表示一个客户信息
type Customer struct {
	Id     int
	Name   string
	Gander string
	Age    int
	Phone  string
	Email  string
}

//编写一个工厂模式，返回一个Customer的实例
func NewCustomer(id int, name string, gander string,
	age int, phone string, email string) *Customer {
	return &Customer{
		Id:     id,
		Name:   name,
		Gander: gander,
		Age:    age,
		Phone:  phone,
		Email:  email,
	}
}

func (c Customer) GetInfo() string {
	return fmt.Sprintf("%v\t%v\t%v\t%v\t%v\t%v",
		c.Id, c.Name, c.Gander, c.Age, c.Phone, c.Email)
}
