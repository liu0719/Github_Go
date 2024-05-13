package processes

import (
	"chatroom/client/utils"
	"chatroom/common/message"
	"encoding/json"
	"fmt"
	"net"
	"strconv"
)

type UserProcerss struct {
	//暂时不需要字段
	User *message.User
}

func (this *UserProcerss) Login(userId int, userPwd string) (err error) {
	c, err := net.Dial("tcp", "192.168.1.3:8888")
	if err != nil {
		fmt.Println("连接服务器失败:", err)
		return
	}
	defer c.Close()
	var mes message.Message
	mes.Type = message.LoginMesType
	var loginMes message.LoginMes
	loginMes.UserId = userId
	loginMes.UserPwd = userPwd

	data, err := json.Marshal(loginMes)
	if err != nil {
		fmt.Println("序列化失败:", err)
		return
	}

	mes.Data = string(data)

	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("序列化失败:", err)
		return
	}

	tf := &utils.Transfer{
		C: c,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println(err)
		return
	}

	//处理服务器返回的消息
	mes, err = tf.ReadPkg()
	if err != nil {
		fmt.Println(err)
		return
	}
	var loginResMes message.LoginResMes
	err = json.Unmarshal([]byte(mes.Data), &loginResMes)
	if err != nil {
		fmt.Println(err)
		return
	}
	if loginResMes.Code == 200 {
		//初始化CurUser
		CurUser.C = c
		CurUser.UserId = userId
		CurUser.UserName = loginResMes.UserName
		CurUser.UserStatus = message.UserOnline
		fmt.Println("登陆成功,", loginResMes.UserName)
		fmt.Println()
		//jieshouloginResMes中的Users数组，遍历存入本地数组
		for _, v := range loginResMes.Users {
			if v.UserId == userId {
				continue
			}
			user := &message.User{
				UserId:     v.UserId,
				UserName:   v.UserName,
				UserStatus: message.UserOnline,
			}
			onlineUsers[v.UserId] = user
		}
		go ServerProcessMes(c)
		for {
			ShowMenu(&loginResMes.User)
		}
	} else {
		fmt.Printf("\n%v\n", loginResMes.Error)
	}
	return
}

func (this *UserProcerss) Register(useId int,
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
	err = tf.WritePkg(data)
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
	return
}

//用户离线
func (this *UserProcerss) Offline(user *message.User) {

	c, err := net.Dial("tcp", "192.168.1.3:8888")
	if err != nil {
		fmt.Printf("\n连接服务器失败,可能未开启服务\n\n")
		return
	}
	defer c.Close()

	var mes message.Message
	mes.Type = message.NotifyUserStatusMesType

	var notifyUserStatusMes message.NotifyUserStatusMes
	notifyUserStatusMes.UserId = user.UserId
	notifyUserStatusMes.UserName = user.UserName
	notifyUserStatusMes.Status = user.UserStatus

	data, err := json.Marshal(notifyUserStatusMes)
	if err != nil {
		fmt.Println(err)
		return
	}

	mes.Data = string(data)

	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println(err)
		return
	}

	tf := &utils.Transfer{
		C: c,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func (this *UserProcerss) SendList(user *message.User) {
	fmt.Println("列表如下(请输入ID选择用户发送消息):")
	for i, v := range onlineUsers {
		var userstatus string
		switch v.UserStatus {
		case 0:
			userstatus = "在线"
		case 1:
			userstatus = "离开"
		case 2:
			userstatus = "忙碌"
		default:
			userstatus = "未知"
		}
		fmt.Printf("用户:%v(ID:%v)\t状态:%v\n", v.UserName, i, userstatus)
	}
	fmt.Println()
	var toUserIdStr string
	var sendContent string
	fmt.Println("请输入ID:")
	fmt.Scanln(&toUserIdStr)
	toUserId, err := strconv.Atoi(toUserIdStr)
	if err != nil {
		fmt.Println("输入有误,请重试")
		fmt.Println()
		return
	}
	var loop bool = true
	for id := range onlineUsers {
		if id == toUserId {
			loop = false
			break
		}
	}
	if loop {
		fmt.Println("暂未找到该ID,请重新输入")
		fmt.Println()
		return
	}
	fmt.Println("已选择该用户，请输入你发送的内容:")
	fmt.Scanln(&sendContent)
	sp := &SmsProcess{}
	err = sp.SendPerson(user, toUserId, sendContent)
	if err != nil {
		fmt.Println("发送失败,未知错误,请稍后重试")
		return
	}
	fmt.Println("发送成功")
}
