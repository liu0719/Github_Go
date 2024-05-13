package main

import (
	"context"
	"fmt"
	"go_micro/grpc_TLS/server/proto/greeter"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type Hello struct {
}

func (Hello) SayHello(c context.Context, req *greeter.HelloReq) (*greeter.HelloRes, error) {
	fmt.Println(req.Name)
	return &greeter.HelloRes{
		Message: "你好" + req.GetName(),
	}, nil
}
func main() {
	//服务加密
	// 传入文件或者cerd进行加密
	cred, err := credentials.NewServerTLSFromFile("E:\\go\\src\\go_micro\\grpc_TLS\\server\\key\\test.pem", "E:\\go\\src\\go_micro\\grpc_TLS\\server\\key\\test.key")
	if err != nil {
		fmt.Println(err)
	}
	//获取到的加密可以传入新建cert中
	grpcServer := grpc.NewServer(grpc.Creds(cred))
	greeter.RegisterGreeterServer(grpcServer, &Hello{})
	listener, err := net.Listen("tcp", "10.50.142.216:80")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("监听已开始，正在等待连接")
	err = grpcServer.Serve(listener)
	if err != nil {
		fmt.Println(err)
	}
}
