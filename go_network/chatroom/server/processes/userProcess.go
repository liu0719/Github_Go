package processes

import (
	"encoding/json"
	"fmt"
	message "go_network/chatroom/common/messgae"
	"go_network/chatroom/server/model"
	"go_network/chatroom/server/utils"
	"net"
)

type UserProcess struct {
	//分析字段
	C net.Conn
	//添加一个字段，表示是哪个用户的
	userId int
}

//编写通知所有用户的方法
func (this *UserProcess) NotifyOthersOnlineUser() {
	//上线的的Id
	// 遍历onlineUser,然后一个一个的发送NotifyUserStatusMesthis
	for id, up := range userMgr.onlineUsers {
		if id == this.userId {
			continue
		}
		//开始通知[单独方法]
		up.NotifyMeOnline(this.userId)
	}
	fmt.Println("已向该用户推送最新在线列表")
}

func (this *UserProcess) NotifyMeOnline(userId int) {
	//组装我们要返回的消息
	var mes message.Message
	mes.Type = message.NotifyUserStatusMesType

	//创建NotifyUserStatusMes
	var notifyUserStatusMes message.NotifyUserStatusMes
	notifyUserStatusMes.UserId = userId
	notifyUserStatusMes.Status = message.UserOnline

	//赋值完成后就转为json序列化
	data, err := json.Marshal(notifyUserStatusMes)
	if err != nil {
		fmt.Println("序列化失败:", err)
	}
	//赋给mes.data
	mes.Data = string(data)

	//再次对mes进行序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("序列化失败:", err)
	}

	//创建tranfer实例
	tf := utils.Transfer{
		C: this.C,
	}
	err = tf.WriterPkg(data)
	if err != nil {
		fmt.Println("NotifyMeOnline err:", err)
		return
	}
}

func (user *UserProcess) ServerProcessRegister(mes *message.Message) (err error) {
	//核心处理代码
	// 1.先从mes中取mes.Data,并直接反序列化成RegisterMes
	var registerMes message.RegisterMes
	err = json.Unmarshal([]byte(mes.Data), &registerMes)
	if err != nil {
		fmt.Println(err)
		return
	}

	//要先声明一个返回的的resMes
	var resMes message.Message
	resMes.Type = message.RegisterResMesType

	//在声明一个LoginResMes
	var registerResMes message.RegisterResMes

	// 这里需要到数据库完成注册·
	// 1、使用model.MyUserDao到Redis去验证
	err = model.MyUserDao.Register(&registerMes.User)
	if err != nil {
		if err == model.Error_USERR_EXISTS {
			registerResMes.Code = 505
			registerResMes.Error = err.Error()
		} else {
			registerResMes.Code = 506
			registerResMes.Error = "服务器正在升级..."
		}
	} else {
		registerResMes.Code = 200
	}
	// 3.RegisterResMes反序列化
	data, err := json.Marshal(registerResMes)
	if err != nil {
		fmt.Println("json.Marshal fail", err)
	}
	resMes.Data = string(data)
	//4.对resMes序列化
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.Marshal fail", err)
	}

	//因为分层(mvc)我们先传建一个Transfer实例，然后读取
	tf := &utils.Transfer{
		C: user.C,
	}
	//发送封装到writerPkg中
	err = tf.WriterPkg(data)

	return
}

//编写一个函数serverProcessLogin函数，处理登录请求
func (this *UserProcess) ServerProcessLogin(mes *message.Message) (err error) {
	//核心处理代码
	// 1.先从mes中取mes.Data,并直接反序列化成LoginMes

	var loginMes message.LoginMes
	err = json.Unmarshal([]byte(mes.Data), &loginMes)

	if err != nil {
		fmt.Println(err)
		return
	}

	//要先声明一个返回的的resMes
	var resMes message.Message
	resMes.Type = message.LoginResMesType

	//在声明一个LoginResMes
	var loginResMes message.LoginResMes

	// 这里需要到数据库完成验证

	// 1、s使用model.MyUserDao到Redis去验证
	user1, err := model.MyUserDao.Login(loginMes.UserId, loginMes.UserPwd)
	if err != nil {
		if err == model.Error_USERR_NOTEXISTS {
			loginResMes.Code = 500
			loginResMes.Error = err.Error()
		} else if err == model.Error_USER_PWD {
			loginResMes.Code = 403
			loginResMes.Error = err.Error()
		} else {
			loginResMes.Code = 505
			loginResMes.Error = "服务器正在升级..."
		}

		//2.这里我们先测试成功，然后可以根据具体返回错误信息
	} else {
		loginResMes.Code = 200
		loginResMes.UserName = user1.UserName
		fmt.Printf("ID:%v,密码:%v,昵称:%v,状态:%v,已登录\n",
			user1.UserId, user1.UserPwd, user1.UserName, user1.UserStatus)
		// 将登陆的用户的userId,赋给this
		this.userId = loginMes.UserId
		//这里因为登陆成功，要讲该用户添加到在线列表中
		userMgr.AddOnlineUser(this)
		this.NotifyOthersOnlineUser()
		for id := range userMgr.onlineUsers {
			loginResMes.UsersId = append(loginResMes.UsersId, id)
		}
	}
	// //如果Id=100密码=123则合法
	// if loginMes.UserId == 100 && loginMes.UserPwd == "123" {
	// 	//合法
	// 	loginResMes.Code = 200
	// } else {
	// 	//不合法
	// 	loginResMes.Code = 500
	// 	loginResMes.Error = "账号或密码错误"
	// }
	// 3.loginResMes反序列化
	data, err := json.Marshal(loginResMes)
	if err != nil {
		fmt.Println("json.Marshal fail", err)
	}
	resMes.Data = string(data)

	//4.对resMes序列化
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.Marshal fail", err)
	}

	//因为分层(mvc)我们先传建一个Transfer实例，然后读取
	tf := &utils.Transfer{
		C: this.C,
	}
	//发送封装到writerPkg中
	err = tf.WriterPkg(data)
	return
}
