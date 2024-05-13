package main

import (
	"fmt"
	message "go_network/chatroom/common/messgae"
	process "go_network/chatroom/server/processes"
	"go_network/chatroom/server/utils"
	"io"
	"net"
)

//创建一个Processor的结构体
type Processor struct {
	C net.Conn
}

//编写一个SertverProcessMes函数
//根据客户端发送种类不同，决定调用哪个函数
func (pro *Processor) serverProcessMes(mes *message.Message) (err error) {
	//验证服务器是否收到发送的消息

	fmt.Println("mes=", mes)

	switch mes.Type {
	case message.LoginMesType:
		//处理登录信息
		//创建一个UserProcess实例
		up := &process.UserProcess{
			C: pro.C,
		}
		err = up.ServerProcessLogin(mes)
		if err != nil {
			fmt.Println(err)
			return
		}
	case message.RegisterMesType:
		//处理注册信息
		up := &process.UserProcess{
			C: pro.C,
		}
		err = up.ServerProcessRegister(mes)
		if err != nil {
			fmt.Println(err)
			return
		}
	case message.SmsMesType:
		//发送消息处理
		//创建一个SmsProcess实例完成转发世界消息
		sm := &process.SmsProcess{}
		sm.SendGroupMes(mes)
	default:
		fmt.Println(pro.C.RemoteAddr().String(), ":发送的消息类型不存在,无法处理")
	}
	return
}
func (pro *Processor) process() (err error) {
	// 循环读取客户端发送的消息
	for {
		//读取直接封装成一个函数readPkg()返回Message error
		tf := &utils.Transfer{
			C: pro.C,
		}
		mes, err := tf.ReadPkg()
		if err != nil {
			if err == io.EOF {
				fmt.Println("客服端已退出...")
			}
			fmt.Println(err)
			return err
		}
		err = pro.serverProcessMes(&mes)
		if err != nil {
			fmt.Println(err)
		}
	}
}
