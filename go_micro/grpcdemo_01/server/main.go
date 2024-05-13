package main

import (
	"context"
	"fmt"
	"go_micro/grpcdemo_01/server/greeter"
	"net"

	"google.golang.org/grpc"
)

type Hello struct {
}

func (Hello) SayHello(c context.Context, req *greeter.HelloReq) (*greeter.HelloRes, error) {
	fmt.Println(req.Name + "连接了服务")
	fmt.Println(c)
	return &greeter.HelloRes{
		Message: "你好" + req.GetName(),
	}, nil
}
func main() {
	grpcServer := grpc.NewServer()

	greeter.RegisterGreeterServer(grpcServer, &Hello{})

	listener, err := net.Listen("tcp", "10.50.142.216:80")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("服务已经启动...")
	err = grpcServer.Serve(listener)
	if err != nil {
		fmt.Println(err)
	}
}
