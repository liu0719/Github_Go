package main

import (
	"fmt"
	"net"
	"net/rpc"
)

type Goods struct {
}
type AddGoods struct {
	Id      int
	Tital   string
	Author  string
	Context string
}
type AddGoodsRes struct {
	Success bool
	Msg     string
}

func (Goods) AddGoods(req AddGoods, res *AddGoodsRes) error {
	fmt.Println(req)
	*res = AddGoodsRes{
		Success: true,
		Msg:     "哈哈哈",
	}
	return nil
}
func main() {
	err := rpc.RegisterName("goods", new(Goods))
	if err != nil {
		fmt.Println(err)
	}
	listener, err := net.Listen("tcp", "10.50.142.216:80")
	if err != nil {
		fmt.Println(err)
	}
	defer listener.Close()
	for {
		fmt.Println("监听端口成功,正在等待连接...")
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
		}
		rpc.ServeConn(conn)

	}
}
