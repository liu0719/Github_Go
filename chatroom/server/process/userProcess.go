package process

import (
	"chatroom/common/message"
	"chatroom/server/model"
	"chatroom/server/utils"
	"encoding/json"
	"fmt"
	"net"
)

type UserProcess struct {
	C net.Conn
	//添加一个字段，表明是那个用户由于维护在线用户
	model.User
}

//编写通知所有用户的方法
func (this *UserProcess) NotifyOthersOnlineUser() {
	//上线的的Id
	// 遍历onlineUser,然后一个一个的发送NotifyUserStatusMesthis
	for id, up := range userMgr.onlineUsers {
		fmt.Println(id, up)
		if id == this.UserId {
			continue
		}
		//开始通知[单独方法]
		up.NotifyMeOnline(this.UserId, this.UserName, this.UserStatus)
	}
	fmt.Println("已向该用户推送最新在线列表")
}

func (this *UserProcess) NotifyMeOnline(userId int, userName string, userStatus int) {
	//组装我们要返回的消息
	var mes message.Message
	mes.Type = message.NotifyUserStatusMesType

	//创建NotifyUserStatusMes
	var notifyUserStatusMes message.NotifyUserStatusMes
	notifyUserStatusMes.UserId = userId
	notifyUserStatusMes.UserName = userName
	notifyUserStatusMes.Status = userStatus

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
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("NotifyMeOnline err:", err)
		return
	}
}

func (this *UserProcess) serverProcessLogin(mes *message.Message) (err error) {

	var loginMes message.LoginMes
	err = json.Unmarshal([]byte(mes.Data), &loginMes)
	if err != nil {
		fmt.Println(err)
		return
	}

	var resMes message.Message
	resMes.Type = message.LoginResMesType

	var loginResMes message.LoginResMes

	//到数据库中验证，调用model层方法
	user, err := model.MyUserDao.Login(loginMes.UserId, loginMes.UserPwd)
	if err != nil {
		if err == model.Error_USER_NOTEXISTS {
			loginResMes.Code = 500
			loginResMes.Error = err.Error()
		} else if err == model.Error_USER_PWD {
			loginResMes.Code = 403
			loginResMes.Error = err.Error()
		} else {
			loginResMes.Code = 505
			loginResMes.Error = "服务器正在升级..."
		}
	} else {
		loginResMes.Code = 200
		loginResMes.UserName = user.UserName
		loginResMes.UserId = user.UserId
		//将登录成功的用户ID赋给this,调用方法将其存入在线用户map
		this.UserId = user.UserId
		this.UserName = user.UserName
		userMgr.AddOnlineUser(this)
		this.NotifyOthersOnlineUser()
		for _, v := range userMgr.onlineUsers {
			loginResMes.Users = append(loginResMes.Users, (*message.User)(&v.User))
		}
		// for id := range userMgr.onlineUsers {
		// 	loginResMes.UserId = append(loginResMes.UserId, id)
		// }
		fmt.Printf("\nID:%v,密码:%v,昵称:%v,状态:%v,已登录\n\n",
			user.UserId, user.UserPwd, user.UserName, user.UserStatus)
	}
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
	err = tf.WritePkg(data)
	return
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
		if err == model.Error_USER_EXISTS {
			registerResMes.Code = 505
			registerResMes.Error = err.Error()
		} else {
			registerResMes.Code = 506
			registerResMes.Error = "服务器正在升级..."
		}
	} else {
		registerResMes.Code = 200
		fmt.Printf("\n新用户:ID:%v,密码:%v,昵称:%v,状态:%v,已注册成功\n\n",
			registerMes.UserId, registerMes.UserPwd, registerMes.UserName, registerMes.UserStatus)
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
	err = tf.WritePkg(data)

	return
}

//处理离线
func (this *UserProcess) ServerProcessOffline(mes *message.Message) (err error) {
	// 1.先从mes中取mes.Data,并直接反序列化成RegisterMes
	var NotifyUserStatusMes message.NotifyUserStatusMes
	err = json.Unmarshal([]byte(mes.Data), &NotifyUserStatusMes)
	if err != nil {
		fmt.Println(err)
		return
	}
	this.UserId = NotifyUserStatusMes.UserId
	this.UserStatus = NotifyUserStatusMes.Status
	this.UserName = NotifyUserStatusMes.UserName
	userMgr.AddOnlineUser(this)
	this.NotifyOthersOnlineUser()
	return
}
