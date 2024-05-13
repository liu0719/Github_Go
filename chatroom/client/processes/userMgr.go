package processes

import (
	"chatroom/client/model"
	"chatroom/common/message"
	"fmt"
)

//客户端维护的map
var onlineUsers map[int]*message.User = make(map[int]*message.User, 10)

var CurUser model.CurUser //我们在用户登录成功后完成对CurUser的初始化

//在客户端显示当前在线的用户
func outputOnlineUser() {
	//遍历客户端维护的map
	fmt.Println("用户列表:")
	for id, user := range onlineUsers {
		var userstatus string
		switch user.UserStatus {
		case 0:
			userstatus = "在线"
		case 1:
			userstatus = "离开"
		case 2:
			userstatus = "忙碌"
		}

		fmt.Printf("用户:%v(ID:%v)状态:%v\n", user.UserName, id, userstatus)
	}
	fmt.Println()
}

//编写一个方法，处理返回的NotifyStatusMes
func UpdateUserStatus(notifyUserStatusMes *message.NotifyUserStatusMes) {

	//优化
	user, ok := onlineUsers[notifyUserStatusMes.UserId]
	if !ok { //原来没有
		user = &message.User{
			UserId:     notifyUserStatusMes.UserId,
			UserName:   notifyUserStatusMes.UserName,
			UserStatus: notifyUserStatusMes.Status,
		}
		onlineUsers[notifyUserStatusMes.UserId] = user
	}
	//如果有了就只修改状态
	user.UserStatus = notifyUserStatusMes.Status
	onlineUsers[notifyUserStatusMes.UserId] = user
	// outputOnlineUser()
}
