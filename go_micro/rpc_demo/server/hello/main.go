package main

import (
	"fmt"
	"net"
	"net/rpc"
)

// 定义一个远程调用的方法
type Hello struct {
}

/*
1.方法只能有两个可序列化的参数，第二个参数是指针类型

	req表示客户端传来的数据
	res表示个客户端返回数据
	可以根据实际需求改变类型，
	但不能是channle complex func 类型，因为不能进行序列化

2.方法返回一个error类型，同时必须是公开的方法
*/
func (Hello) SayHello(req string, res *string) error {

	*res = "你好" + req + "真帅"
	fmt.Println(req, *res)
	return nil
}
func main() {
	// 1.注册RPC服务
	//第一个参数是rpc服务名，第二个参数是绑定的方法，实例化绑定的结构体
	err := rpc.RegisterName("hello", new(Hello))
	if err != nil {
		fmt.Println(err)
	}

	//2.监听服务端某一个端口
	//什么协议，运行的地址
	lister, err := net.Listen("tcp", ":80")
	if err != nil {
		fmt.Println(err)
	}
	//3.应用退出时，关闭监听
	defer lister.Close()
	for {
		fmt.Println("开始建立连接")
		//4.接收客户端的连接
		conn, err := lister.Accept()
		if err != nil {
			fmt.Println(err)
		}

		//4.绑定一个服务
		rpc.ServeConn(conn)
	}

}
