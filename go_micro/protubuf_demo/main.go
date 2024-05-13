package main

import (
	"fmt"
	"go_micro/protubuf_demo/proto/userService"

	"google.golang.org/protobuf/proto"
)

func main() {
	user := &userService.Userinfo{
		Username: "张三",
		Age:      20,
		Hobby:    []string{"吃饭", "睡觉", "打豆豆"},
	}
	// fmt.Println(user)
	//获取某一个属性
	fmt.Println(user.GetUsername())
	// protobuf的序列化，和json类似
	data, _ := proto.Marshal(user)
	fmt.Println(data)
	//反序列化
	var newUser userService.Userinfo
	_ = proto.Unmarshal(data, &newUser)
	//这里打印出来有空的，是因为结构体力还有生成的其他字段
	fmt.Println(newUser.GetHobby())
	fmt.Println("----------")
	fmt.Println(user.GetType())
}
