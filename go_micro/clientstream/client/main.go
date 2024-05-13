package main

import (
	"bufio"
	"context"
	"fmt"
	"go_micro/clientstream/server/proto/greeter"
	"os"

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
	// 普通调用
	res, err := client.SayHello(context.Background(), &greeter.HelloReq{
		Name: "刘广硕",
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res.Message)

	// 调用定义带流的方法，会返回一个流类型的值
	//stream流
	stream, err := client.UpdateHello(context.Background())
	if err != nil {
		fmt.Println(err)
	}
	reader := bufio.NewReader(os.Stdin)
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			fmt.Println(err)
		}
		if string(line) == "close" {
			res, err := stream.CloseAndRecv()
			if err != nil {
				fmt.Println(err)

			}
			fmt.Println(res)
			return

		}
		req := &greeter.HelloReq{
			Name: string(line),
		}
		stream.Send(req)
		fmt.Println("发送成功")
	}
}
