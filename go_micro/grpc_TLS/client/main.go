package main

import (
	"context"
	"fmt"
	"go_micro/grpc_TLS/server/proto/greeter"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	//传入证书pem验证才能调用服务端
	cred, err := credentials.NewClientTLSFromFile("E:\\go\\src\\go_micro\\grpc_TLS\\server\\key\\test.pem", "*.haha.com") //域名，真实的要去浏览器里获取
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(cred)
	conn, err := grpc.Dial("10.50.142.216:80", grpc.WithTransportCredentials(cred))
	if err != nil {
		fmt.Println(err)
	}
	client := greeter.NewGreeterClient(conn)
	res, err := client.SayHello(context.Background(), &greeter.HelloReq{
		Name: "刘广硕",
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res.Message)
}
