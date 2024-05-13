package main

import (
	"fmt"
	"go_network/chatroom/client/processes"
	"os"
	"time"
)

var (
	userId   int
	userPwd  string
	userName string
)

func mainMenu() {
	var key int
	var loop = true
	for loop {
		fmt.Println("-------------欢迎使用多人聊天系统-------------")
		fmt.Println("\t\t 1.登录账号")
		fmt.Println("\t\t 2.注册账号")
		fmt.Println("\t\t 3.退出系统")
		fmt.Println("请选择(1-3):")
		fmt.Scanln(&key)
		switch key {
		case 1:
			fmt.Println("登录账号:")
			//说明用户要登录
			fmt.Println("请输入账号")
			fmt.Scanln(&userId)
			fmt.Println("请输入密码")
			fmt.Scanln(&userPwd)
			//完成登录
			// 1.创建一个userProcess的实例
			up := &processes.UserProcerss{}
			up.Login(userId, userPwd)
			time.Sleep(time.Second * 2)
			// loop = false
		case 2:
			fmt.Println("注册用户:")
			fmt.Println("请输入注册的ID:")
			fmt.Scanln(&userId)
			fmt.Println("请输入注册密码:")
			fmt.Scanln(&userPwd)
			fmt.Println("请输入注册昵称:")
			fmt.Scanln(&userName)
			//2.调用UserProcess,完成注册的请求
			up := &processes.UserProcerss{}
			up.Register(userId, userPwd, userName)
			time.Sleep(time.Second * 2)
			// loop = false
		case 3:
			fmt.Println("退出系统")
			// loop = false
			os.Exit(0)
		default:
			fmt.Println("输入有误，请重试。。。")
		}
	}
}
func main() {
	mainMenu()
}
