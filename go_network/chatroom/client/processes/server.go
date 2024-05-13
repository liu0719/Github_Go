package processes

import (
	"encoding/json"
	"fmt"
	"go_network/chatroom/client/utils"
	message "go_network/chatroom/common/messgae"
	"net"
	"os"
)

//显示登陆成功后的界面。。。
func ShowMenu(userName string) {
	fmt.Println("----------欢迎您,", userName, "-----------")
	fmt.Println("----------1.显示在线用户列表----")
	fmt.Println("----------2.发送消息------------")
	fmt.Println("----------3.信息列表------------")
	fmt.Println("----------4.退出系统------------")
	fmt.Println("请选择(1-4);")
	var key int
	var content string

	//因为我们总会使用到SmsProcess实例,因此我们将其定义在switch外面
	smsProcess := &SmsProcess{}
	fmt.Scanln(&key)
	switch key {
	case 1:
		outputOnlineUser()
	case 2:
		fmt.Println("请输入你要发送的内容(世界频道):")
		fmt.Scanln(&content)
		smsProcess.SendGroup(content, userName)
	case 3:
		fmt.Println("信息列表")
	case 4:
		fmt.Println("退出了系统...")
		os.Exit(0)
	default:
		fmt.Println("输入错误,请重试")
	}
}

//和服务器端保持通讯
func ServerProcessMes(c net.Conn) {
	//创建transfer实例，不停地读取服务器发送的消息
	tf := &utils.Transfer{
		C: c,
	}
	for {
		fmt.Println("正在等待读取服务端发送的消息")
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
			UpdateUserStatus(&notifyUserStatusMes)
			// 2.把这个用户端的信息状态,保存到客户map[int]user中
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
