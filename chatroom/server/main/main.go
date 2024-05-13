package main

import (
	"chatroom/server/model"
	"chatroom/server/process"
	"fmt"
	"net"
)

func mainprocess(c net.Conn) {
	defer c.Close()
	pro := &process.Processor{
		C: c,
	}
	pro.Process()
}
func initUserDao() {
	model.MyUserDao = *model.NewUserDao(process.Pool)
}
func main() {
	//初始化链接池
	process.InitPool("192.168.1.3:6379", 16, 0, 300)
	initUserDao()
	//开始监听
	listen, err := net.Listen("tcp", "192.168.1.3:8888")
	if err != nil {
		fmt.Println("监听端口失败:", err)
	}
	fmt.Println("服务器开始监听...")
	//延时关闭
	defer listen.Close()
	// 循环读取是否有人连接端口
	for {
		c, err := listen.Accept()
		if err != nil {
			fmt.Println("有人连接失败")
		}
		fmt.Println("有人链接成功:", c.RemoteAddr().String())
		go mainprocess(c)
	}

}
