package main

import (
	"fmt"
	"go_network/chatroom/server/model"
	"net"
)

func mainprocess(c net.Conn) {
	//读取客服端发送的消息
	defer c.Close()
	//调用总控
	pro := &Processor{
		C: c,
	}
	err := pro.process()
	if err != nil {
		fmt.Println("携程出错", err)
	}
}

//完成对UserDao1初始化
func initUserDao() {
	//这里的pool本身就是一个全局的变量
	//这里要注意初始化的顺序问题
	model.MyUserDao = *model.NewUserDao(pool)

}
func main() {
	//注意初始化顺序，inituserDao要依赖于initPool
	//服务器启东，初始化链接池
	initPool("192.168.1.3:6379", 16, 0, 300)
	initUserDao()
	listen, err := net.Listen("tcp", "0.0.0.0:8888")
	if err != nil {
		fmt.Println("开启监听失败", err)
	}
	//提示信息
	fmt.Println("服务器在196.168.1.3:8888监听")
	defer listen.Close()
	//监听成功就等待客服端来连接服务器
	for {
		fmt.Println("服务器监听连接中...")
		c, err := listen.Accept()
		if err != nil {
			fmt.Println("与人连接失败:", err)
		}
		go mainprocess(c)
	}
}
