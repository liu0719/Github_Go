package main

import (
	"context"
	"fmt"
	"go_micro/grpc_TLS/server/proto/greeter"

	"google.golang.org/grpc"
	_ "google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

// 自定义token结构体
type ClientTokenAuth struct {
}

// 下面这两个方法实现了grpc提供的接口，服务端就可以获取到token，进行验证
func (ClientTokenAuth) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"appid":  "haha",
		"appkey": "123123",
	}, nil
}

// 要不要开启TSL认证
func (ClientTokenAuth) RequireTransportSecurity() bool {
	// 返回false不用开启TSL认证
	//要加密的话就要把TSL证书打开
	return false
}

func main() {
	//
	// cred, err := credentials.NewClientTLSFromFile("E:/go/src/go_micro/grpc_TLS/server/key/test.pem", "*.haha.com") //域名，真实的要去浏览器里获取
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// 声明一个链接参数的数组
	var opts []grpc.DialOption
	// 把创建的证书加入
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	//把自定义的结构体appent
	opts = append(opts, grpc.WithPerRPCCredentials(ClientTokenAuth{}))
	conn, err := grpc.Dial("10.50.142.216:80", opts...)
	if err != nil {
		fmt.Println(err)
	}
	client := greeter.NewGreeterClient(conn)
	if err != nil {
		fmt.Println(err)
	}
	res, err := client.SayHello(context.Background(), &greeter.HelloReq{
		Name: "刘广硕",
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res.Message)
}
