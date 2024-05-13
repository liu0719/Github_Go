package main

import (
	"context"
	"fmt"
	"go_micro/clientstream/server/proto/greeter"
	"io"
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

func (Hello) UpdateHello(stream greeter.Greeter_UpdateHelloServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			stream.SendAndClose(&greeter.HelloRes{
				Message: "大家好，" + req.GetName(),
			})
			if err != nil {
				fmt.Println(err)
			}

		}

		fmt.Println("收到的值", req)
	}

}
func (Hello) MustEmbedUnimplementedGreeterServer() {}
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
