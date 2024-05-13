package main

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
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

func main() {
	conn, err := net.Dial("tcp", "10.50.142.216:80")
	if err != nil {
		fmt.Println(err)
	}
	//建立一个基于json编解码的rpc服务
	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))
	var reply AddGoodsRes
	//替换conn
	client.Call("goods.AddGoods", AddGoods{
		Id:      10,
		Tital:   "三国演义",
		Author:  "罗贯中",
		Context: "东汉末年",
	}, &reply)
	fmt.Println(reply)
}

/*
把默认的rpc 改为jsonrpc
1.rpc.Dial调换为net.Dial
2.增加建立基于json编解码的rpc服务 cilent:=rpc.NewClientWirterCodec(jsonrpc.NewClientCodec(conn))
3.conn.Call需要改为client.Call
*/
