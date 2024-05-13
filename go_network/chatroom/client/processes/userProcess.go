package processes

import (
	"encoding/json"
	"fmt"
	message "go_network/chatroom/common/messgae"
	"go_network/chatroom/server/utils"
	"net"
	"os"
)

type UserProcerss struct {
	//暂时不需要字段
}

func (user *UserProcerss) Register(userId int,
	userPwd string, userName string) (err error) {
	//1.连接到服务器
	c, err := net.Dial("tcp", "192.168.1.3:8888")
	if err != nil {
		fmt.Printf("\n连接服务器失败,可能未开启服务\n\n")
		return
	}
	defer c.Close()

	// 2.通过conn发送消息给服务器
	var mes message.Message
	mes.Type = message.RegisterMesType

	//3.创建一个loginMes结构体当做发送的Message的Data
	var registerMes message.RegisterMes
	registerMes.User.UserId = userId
	registerMes.User.UserPwd = userPwd
	registerMes.User.UserName = userName

	//4.将registerMes序列化后放进Message.Data中
	data, err := json.Marshal(registerMes)
	if err != nil {
		fmt.Println("发送的数据体序列化失败:", err)
		return
	}
	mes.Data = string(data)

	//5.将mes序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("发送的数据序列化失败:", err)
		return
	}
	//创建Tranfser调用utils包
	tf := &utils.Transfer{
		C: c,
	}
	err = tf.WriterPkg(data)
	if err != nil {
		fmt.Println("注册消息发送失败:", err)
	}

	//这里还需要处理服务断返回的消息
	mes, err = tf.ReadPkg() //mes就是RegisterResMes
	if err != nil {
		fmt.Println("readPkg(c) err:", err)
		return
	}

	//将mes的data部分反序列化RegisterResMes
	var registerResMes message.RegisterResMes
	err = json.Unmarshal([]byte(mes.Data), &registerResMes)
	if err != nil {
		fmt.Println(err)
	}
	if registerResMes.Code == 200 {
		fmt.Println("注册成功,请登录")

		// 这里要在客服端启动一个协程
		//该携程保持和服务器端的通信，如果服务器有数据要退送给客户端
		// 则接收并显示在客户端的终端
	} else {
		fmt.Printf("\n%v\n\n", registerResMes.Error)
	}
	os.Exit(0)
	return
}

//登录
func (this *UserProcerss) Login(userId int, userPwd string) (err error) {
	//1.连接到服务器
	c, err := net.Dial("tcp", "192.168.1.3:8888")
	if err != nil {
		fmt.Printf("\n连接服务器失败,可能未开启服务\n\n")
		return
	}
	defer c.Close()
	// 2.通过conn发送消息给服务器
	var mes message.Message
	mes.Type = message.LoginMesType

	//3.创建一个loginMes结构体当做发送的Message的Data
	var loginMes message.LoginMes
	loginMes.UserId = userId
	loginMes.UserPwd = userPwd

	//4.将loginMs序列化后放进Message.Data中
	data, err := json.Marshal(loginMes)
	if err != nil {
		fmt.Println("发送的数据体序列化失败:", err)
		return
	}
	mes.Data = string(data)

	//5.将mes序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("发送的数据序列化失败:", err)
		return
	}
	//6.这时data就是我们要发送的数据
	//先发送的数据的长度
	/*
		pkgLen := uint32(len(data))
		//将长度转成bytes
		var buf [4]byte
		binary.BigEndian.PutUint32(buf[:4], pkgLen)
		n, err := c.Write(buf[:4])
		if n != 4 || err != nil {
			fmt.Println("发送失败:", err)
			return
		}
		// fmt.Println(n)

		//发送消息本身
		_, err = c.Write(data)
		if err != nil {
			fmt.Println("发送失败:", err)
			return
		}
	*/
	tf := &utils.Transfer{
		C: c,
	}
	err = tf.WriterPkg(data)
	if err != nil {
		fmt.Println(err)
	}
	//这里还需要处理服务断返回的消息
	mes, err = tf.ReadPkg()
	if err != nil {
		fmt.Println(err)
	}

	//将mes的data部分反序列化LoginResMes
	var loginResMes message.LoginResMes
	err = json.Unmarshal([]byte(mes.Data), &loginResMes)
	if err != nil {
		fmt.Println(err)
	}
	if loginResMes.Code == 200 {
		//初始化CurUser
		CurUser.C = c
		CurUser.UserId = userId
		CurUser.UserStatus = message.UserOnline
		fmt.Println("登陆成功")

		//可以显示当前用在线列表，遍历loginResMes.UserId
		// fmt.Println("当前在线用户列表如下:")
		for _, v := range loginResMes.UsersId {

			// //不显示自己的id
			//在本地客户端不加入自己ID
			if v == userId {
				continue
			}
			// fmt.Println("用户id:\t", v)

			//完成客户端的onlineUser的初始化
			user := &message.User{
				UserId:     v,
				UserStatus: message.UserOnline,
			}
			onlineUsers[v] = user
		}
		fmt.Println()
		// 这里要在客服端启动一个协程
		//该携程保持和服务器端的通信，如果服务器有数据要退送给客户端
		// 则接收并显示在客户端的终端
		go ServerProcessMes(c)
		// 1.显示登录成功的菜单
		for {
			ShowMenu(loginResMes.UserName)
		}
	} else {
		fmt.Printf("\n%v\n\n", loginResMes.Error)
	}

	return
}
