//负责界面显示
package main

import (
	"fmt"
	"go_base/object_oriented/CustomermManagementSystem/service"
	"time"
)

type customerView struct {
	key             string //接收用户输入
	loop            bool   //表示是否循环的显示主菜单
	customerService *service.CustomerService
}

//显示主菜单
func (c *customerView) mainMenu() {
	for {
		time.Sleep(time.Second)
		fmt.Println("-------------客户信息管理软件-------------")
		fmt.Println("		1.添加客户")
		fmt.Println("		2.修改客户")
		fmt.Println("		3.删除客户")
		fmt.Println("		4.客户列表")
		fmt.Println("		5.退   出")
		fmt.Print("请选择(1-5):")
		fmt.Scanln(&c.key)
		switch c.key {
		case "1":
			c.add()
		case "2":
			c.updata()
		case "3":
			c.delete()
		case "4":
			c.list()
		case "5":
			c.exit()
		default:
			{
				fmt.Println("输入不符合，请重试...")
			}
		}
		if c.loop {
			fmt.Println("退出成功，欢迎下次使用！")
			break
		}

	}
}

//客户列表
func (c *customerView) list() {
	//首先获取到当前的客户信息(在切片中)
	customers := c.customerService.List()
	//显示
	fmt.Println("-----------------客户列表-----------------")
	fmt.Println("编号\t姓名\t性别\t年龄\t电话\t邮箱")
	for _, v := range customers {
		fmt.Println(v.GetInfo())
	}
	fmt.Println()
	fmt.Println("客户列表显示完成,正在召唤主菜单...")
}

//添加用户
//得到用户的输入，构建一个新的customer交给Service(调用sevice中的Add)
func (c *customerView) add() {
	fmt.Println("-----------------添加客户-----------------")
	fmt.Println("姓名:")
	name := ""
	fmt.Scanln(&name)
	fmt.Println("性别:")
	gender := ""
	fmt.Scanln(&gender)
	fmt.Println("年龄:")
	age := 0
	fmt.Scanln(&age)
	fmt.Println("电话:")
	phone := ""
	fmt.Scanln(&phone)
	fmt.Println("邮箱:")
	email := ""
	fmt.Scanln(&email)
	flag := c.customerService.Add(name, gender, age, phone, email)
	if flag {
		fmt.Println("添加客户成功...")
		fmt.Println("该客户编号为", c.customerService.CustNum, "\n请牢记或保存...")
	} else {
		fmt.Println("添加客户失败,请重试")
	}
	fmt.Println()
	time.Sleep(time.Second)
	fmt.Println("正在返回主菜单...")
}

//删除客户
func (c *customerView) delete() {
	fmt.Println("-----------------删除客户-----------------")
	fmt.Println("请输入你要删除的ID(不输入自动放弃):")
	id := -1
	fmt.Scanln(&id)
	if id == -1 {
		fmt.Println("已自动放弃,")
		return //放弃删除
	}
	//调用service里的find方法
	str := c.customerService.FindId(id)
	if str == "" {
		fmt.Println("没有找到该用户，请检查输入或重试")
	} else {
		fmt.Println("找到用户数据如下:")

		fmt.Println(str)
		fmt.Println("是否确认删除?(y/n)默认不删除")
		conform := ""
		fmt.Scanln(&conform)
		if conform == "y" || conform == "Y" {
			flag := c.customerService.Delete(id)
			if flag {
				fmt.Println("删除成功...")
				time.Sleep(time.Second)
			} else {
				fmt.Println("删除失败,请重试...")
			}
		} else if conform == "n" || conform == "N" {
			fmt.Println("取消删除成功...")
		} else {
			fmt.Println("输入错误，请重试...")
		}
	}
	time.Sleep(time.Second)
	fmt.Println("正在返回主菜单...")
}

//修改用户
func (c *customerView) updata() {
	fmt.Println("-----------------修改客户-----------------")
	fmt.Println("请输入你要修改的客户ID(不输入自动放弃):")
	id := -1
	fmt.Scanln(&id)
	if id == -1 {
		fmt.Println("已自动放弃修改,")
		return //放弃删除
	}
	//调用service里的find方法
	str := c.customerService.FindId(id)
	if str == "" {
		fmt.Println("没有找到该用户，请检查输入或重试")
	} else {
		fmt.Println("找到用户数据如下:")
		fmt.Println(str)
		fmt.Println()
		fmt.Println("请输入你新修改的客户信息:")
		fmt.Println("姓名:")
		name := ""
		fmt.Scanln(&name)
		fmt.Println("性别:")
		gender := ""
		fmt.Scanln(&gender)
		fmt.Println("年龄:")
		age := 0
		fmt.Scanln(&age)
		fmt.Println("电话:")
		phone := ""
		fmt.Scanln(&phone)
		fmt.Println("邮箱:")
		email := ""
		fmt.Scanln(&email)
		fmt.Println("是否确认修改?(y/n)默认不修改")
		conform := ""
		fmt.Scanln(&conform)
		if conform == "y" || conform == "Y" {
			flag := c.customerService.UpData(id, name, gender, age, phone, email)
			if flag {
				fmt.Println("修改成功...")
				time.Sleep(time.Second)
			} else {
				fmt.Println("修改失败,请重试...")
			}
		} else if conform == "n" || conform == "N" {
			fmt.Println("取消修改成功...")
		} else {
			fmt.Println("输入错误，请重试...")
		}
	}
	time.Sleep(time.Second)
	fmt.Println("正在返回主菜单...")
}

//退出
func (c *customerView) exit() {
	fmt.Println("确定退出吗(y/n)")
	ex := ""
	fmt.Scanln(&ex)
	if ex == "y" || ex == "Y" {
		c.loop = true
	} else if ex == "n" || ex == "N" {
		fmt.Println("取消退出成功")
	} else {
		fmt.Println("输入错误，请重试...")
	}
}

func main() {
	c := customerView{
		key:  "",
		loop: false,
		//这里完成对customerView结构体的customerService字段的初始化
		customerService: service.NewCustomerService(),
	}

	c.mainMenu()
}
