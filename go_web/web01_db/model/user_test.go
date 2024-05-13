package model

import (
	"fmt"
	"testing"
)

//TestMain可以在测试函数执行之前做一些其他操作
func TestMain(m *testing.M) {
	fmt.Println("测试开始:-------------")
	m.Run()
}
func TestUser(t *testing.T) {
	fmt.Println("开始测试User中的相关方法")
	t.Run("测试添加用户", testAddUser)
	// t.Run("测试获取用户", testGetUserById)
	t.Run("测试获取所有用户", testGetAllUser)
}
func testAddUser(t *testing.T) {
	fmt.Println("测试添加用户")
	user := &Users{}
	user.AddUser()
	user.AddUser2()
}
func testGetUserById(t *testing.T) {
	fmt.Println("测试查询一条记录:")
	user := Users{
		Id: 1,
	}
	//调用
	u, _ := user.GetUserById()
	fmt.Println("得到的服务信息", u)
}

//测试查询所有的User
func testGetAllUser(t *testing.T) {
	fmt.Println("测试查询所有的记录")
	user := &Users{}
	users, _ := user.GetAllUser()
	//便利切片
	for i, v := range users {
		fmt.Printf("第%v个用户=ID:%v,用户名:%v,密码:%v,邮件%v\n",
			i+1, v.Id, v.Username, v.Password, v.Email)
	}
}
