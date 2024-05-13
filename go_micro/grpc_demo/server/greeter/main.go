package main

import (
	"context"
	"fmt"
	"go_micro/grpc_demo/server/greeter/proto/greeter"
	"net"

	"google.golang.org/grpc"
)

type Hello struct {
}

func (Hello) SayHello(c context.Context, req *greeter.HelloReq) (*greeter.HelloRes, error) {
	fmt.Println(req)
	return &greeter.HelloRes{
		Message: "你好" + req.GetName(),
	}, nil
}
func main() {
	//1.初始化grpc对象
	grpcServer := grpc.NewServer()
	// 2.注册服务
	greeter.RegisterGreeterServer(grpcServer, &Hello{})
	// 3.设置监听，指定Ip port
	listener, err := net.Listen("tcp", "localhost:80")
	if err != nil {
		fmt.Println(err)
	}
	defer listener.Close()
	// 4.启动服务
	fmt.Println("服务已经启动...")
	err = grpcServer.Serve(listener)
	if err != nil {
		fmt.Println(err)
	}

}
