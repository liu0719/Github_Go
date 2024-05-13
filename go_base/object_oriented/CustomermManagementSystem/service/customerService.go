//负责业务逻辑
package service

import (
	"fmt"
	"go_base/object_oriented/CustomermManagementSystem/model"
)

//该CustomerService,完成对Customer的操作
//包括增删改查
type CustomerService struct {
	customers []model.Customer
	//盛明明一个字段，表示当前客户的总数
	//该字段后面还可以作为新客户的id+1
	CustNum int
}

//编写一个函数，返回*CustomerService
func NewCustomerService() *CustomerService {
	//为了看到客户在切片中，初始化一个客户
	customerService := &CustomerService{}
	customerService.CustNum = 1
	customer := model.NewCustomer(1, "张三", "男", 20, "122", "zs@sohu.com")
	customerService.customers = append(customerService.customers, *customer)
	return customerService
}

//查看客户列表
func (c *CustomerService) List() []model.Customer {
	return c.customers
}

//添加客户
func (c *CustomerService) Add(name string, gander string,
	age int, phone string, email string) bool {
	customer := model.NewCustomer(c.CustNum+1, name,
		gander, age, phone, email)
	c.customers = append(c.customers, *customer)
	c.CustNum++
	return true
}

//删除客户
func (c *CustomerService) Delete(id int) bool {
	index := 0
	for i, v := range c.customers {
		if id == v.Id {
			index = i
		}
	}
	c.customers = append(c.customers[:index], c.customers[index+1:]...)
	return true
}

//修改客户
func (c *CustomerService) UpData(id int, name string, gander string,
	age int, phone string, email string) bool {
	index := 0
	for i, v := range c.customers {
		if id == v.Id {
			index = i
		}
	}
	c.customers[index].Name = name
	c.customers[index].Age = age
	c.customers[index].Gander = gander
	c.customers[index].Phone = phone
	c.customers[index].Email = email
	return true
}

//删除,修改时找Id
func (c *CustomerService) FindId(id int) (str string) {
	for _, v := range c.customers {
		if id == v.Id {
			str = fmt.Sprintf("ID:%v,姓名:%v,性别:%v,年龄:%v,电话:%v,Email:%v",
				v.Id, v.Name, v.Gander, v.Age, v.Phone, v.Email)
		}
	}
	return str
}
