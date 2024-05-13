package main

import (
	"context"
	"fmt"
	"go_micro/grpc_demo/server/greeter/proto/greeter"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// 1.连接服务器
	/*
		credentials.NewclientTLSFromFile:从输入的证书文件中为客户端构造TLS凭证。
		grpc.withTransportcredentials:配置连接级别的安全R证（例，TLS/SSL)，返回一个DialOption，用于连接服务器。

	*/
	conn, err := grpc.Dial("10.50.142.216:80", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println(err)
	}
	// 2.注册客服端
	client := greeter.NewGreeterClient(conn)

	res, err := client.SayHello(context.Background(), &greeter.HelloReq{
		Name: "刘广硕",
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res.Message)
}
