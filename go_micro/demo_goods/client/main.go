package main

import (
	"fmt"
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

func main() {
	conn, err := rpc.Dial("tcp", "10.50.142.216:80")
	if err != nil {
		fmt.Println(err)
	}
	var reply AddGoodsRes
	conn.Call("goods.AddGoods", AddGoods{
		Id:      10,
		Tital:   "三国演义",
		Author:  "罗贯中",
		Context: "东汉末年",
	}, &reply)
	fmt.Println(reply)
}
