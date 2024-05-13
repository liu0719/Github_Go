package processes

import (
	"chatroom/client/utils"
	"chatroom/common/message"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"time"
)

var (
	userId   int
	userPwd  string
	userName string
)

//登录前主菜单
func MainMenu() {
	var key int
	var loop = true
	for loop {
		fmt.Println("-------------欢迎使用多人聊天系统-------------")
		fmt.Println("\t\t 1.登录账号")
		fmt.Println("\t\t 2.注册账号")
		fmt.Println("\t\t 3.退出系统")
		fmt.Println("请选择(1-3):")
		fmt.Scanln(&key)
		switch key {
		case 1:
			fmt.Println("登录账号:")
			//说明用户要登录
			fmt.Println("请输入账号")
			fmt.Scanln(&userId)
			fmt.Println("请输入密码")
			fmt.Scanln(&userPwd)
			//完成登录
			// 1.创建一个userProcess的实例
			up := UserProcerss{}
			up.Login(userId, userPwd)
			// loop = false
		case 2:
			fmt.Println("注册用户:")
			fmt.Println("请输入注册的ID:")
			fmt.Scanln(&userId)
			fmt.Println("请输入注册密码:")
			fmt.Scanln(&userPwd)
			fmt.Println("请输入注册昵称:")
			fmt.Scanln(&userName)
			up := UserProcerss{}
			up.Register(userId, userPwd, userName)
			// loop = false
		case 3:
			fmt.Println("退出了系统...")
			os.Exit(0)
			// loop = false
		default:
			fmt.Println("输入有误，请重试。。。")
		}
	}

}

//显示登陆成功后的界面。。。
func ShowMenu(user *message.User) {
	fmt.Println("----------欢迎您,", user.UserName, "-----------")
	fmt.Println("----------1.显示用户列表--------")
	fmt.Println("----------2.发送消息(世界)-------")
	fmt.Println("----------3.信息列表(私信)-------")
	fmt.Println("----------4.退出系统------------")
	fmt.Println("请选择(1-4);")
	var key int
	var content string

	//因为我们总会使用到SmsProcess实例,因此我们将其定义在switch外面
	// smsProcess := &SmsProcess{}
	fmt.Scanln(&key)
	switch key {
	case 1:
		outputOnlineUser()
	case 2:
		fmt.Println("请输入你要发送的内容(世界频道):")
		fmt.Scanln(&content)
		sp := SmsProcess{}
		sp.SendGroup(content, user.UserName)
	case 3:
		fmt.Println("信息列表")
		up := &UserProcerss{}
		up.SendList(user)
	case 4:
		up := &UserProcerss{}
		user.UserStatus = message.UserOffline
		up.Offline(user)
		fmt.Println("正在退出系统")
		time.Sleep(time.Second * 2)
		fmt.Println("退出成功,欢迎您下次使用！")
		os.Exit(0)
	default:
		fmt.Println("输入有误，请重试。。。")
	}
}

//和服务器端保持通讯
func ServerProcessMes(c net.Conn) {
	//创建transfer实例，不停地读取服务器发送的消息
	tf := &utils.Transfer{
		C: c,
	}
	for {
		mes, err := tf.ReadPkg()
		if err != nil {
			fmt.Println("read err", err)
			return
		}
		//如果读取到消息，有事下一步处理结果
		switch mes.Type {
		case message.NotifyUserStatusMesType:
			//通知有人上线
			// 1.取出NotifyStatusmes
			var notifyUserStatusMes message.NotifyUserStatusMes
			err = json.Unmarshal([]byte(mes.Data), &notifyUserStatusMes)
			if err != nil {
				fmt.Println("消息信息fan序列化失败:", err)
			}
			fmt.Println("在线列表有更新,按\"1\"查看")
			fmt.Println()
			// 2.把这个用户端的信息状态,保存到客户map[int]user中
			UpdateUserStatus(&notifyUserStatusMes)
		case message.SmsMesType:
			//有人群发消息
			fmt.Println("新的世界消息:")
			outputGroup(&mes)
		default:
			fmt.Println("服务器返回了未知的类型。。。")
		}
		// fmt.Println("mes=", mes)
	}
}
