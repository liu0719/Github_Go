package main

import (
	"fmt"
	"net/rpc"
)

func main() {
	// 1.用rpc连接服务器---Dial()
	conn, err := rpc.Dial("tcp", "10.50.129.238:80")
	if err != nil {
		fmt.Println(err)
	}
	//客户端退出时关闭连接
	defer conn.Close()

	//调用远程函数

	var reply string
	var message string
	fmt.Println("请输入你要发送的信息")
	fmt.Scanln(&message)
	err = conn.Call("hello.SayHello", message, &reply)
	if err != nil {
		fmt.Println(err)
	}
	//获取远程服务（微服务）的数据
	fmt.Println(reply)
}
