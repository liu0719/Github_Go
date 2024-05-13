package main

import (
	"context"
	"fmt"
	"go_micro/grpcdemo_01/server/greeter"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("10.50.142.216:80", grpc.WithTransportCredentials(insecure.NewCredentials()))
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
