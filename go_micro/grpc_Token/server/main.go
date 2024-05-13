package main

import (
	"context"
	"errors"
	"fmt"
	"go_micro/grpc_TLS/server/proto/greeter"
	"net"

	"google.golang.org/grpc"
	_ "google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

type Hello struct {
}

// 业务
func (Hello) SayHello(c context.Context, req *greeter.HelloReq) (*greeter.HelloRes, error) {
	// 业务之前都有拦截器
	md, ok := metadata.FromIncomingContext(c)
	if !ok {
		return nil, errors.New("没有传输token")
	}
	var appId string
	if v, ok := md["appid"]; ok {
		appId = v[0]
	}
	var appKey string
	if v, ok := md["appkey"]; ok {
		appKey = v[0]
	}
	//判断条件，一般条件在数据库中获取得到
	if appId != "haha" || appKey != "123123" {
		return nil, errors.New("验证失败,没有该token")
	}
	//业务处理
	fmt.Println(req.Name)
	return &greeter.HelloRes{
		Message: "你好" + req.GetName(),
	}, nil
}

func main() {
	/*
		//TLS服务加密
		// 传入文件或者cerd进行加密
		cred, err := credentials.NewServerTLSFromFile("E:/go/src/go_micro/grpc_TLS/server/key/test.pem", "E:/go/src/go_micro/grpc_TLS/server/key/test.key")
		if err != nil {
			fmt.Println(err)
		}
		//获取到的加密可以传入新建cert中
		grpcServer := grpc.NewServer(grpc.Creds(cred))
	*/
	// 自定义Token加密
	grpcServer := grpc.NewServer(grpc.Creds(insecure.NewCredentials()))
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
