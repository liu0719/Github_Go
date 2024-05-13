package process

import (
	"chatroom/client/utils"
	"chatroom/common/message"
	"encoding/json"
	"fmt"
	"net"
)

type SmsProcess struct {
	//...暂时不需要字段
}

func (this *SmsProcess) SendGroupMes(mes *message.Message) {
	//遍历服务器端的OnlineUsers map[int]*UserProcess

	//取出mes的内容
	var smsMes message.SmsMes
	err := json.Unmarshal([]byte(mes.Data), &smsMes)
	if err != nil {
		fmt.Println("fan序列化失败:", err)
		return
	}
	fmt.Println(smsMes.UserName)
	data, err := json.Marshal(mes)
	if err != nil {
		fmt.Println("序列化失败,err", err)
		return
	}
	for _, up := range userMgr.onlineUsers {
		//过滤掉自己
		// if id == smsMes.UserId {
		// 	continue
		// }
		this.SendMesToEachOnlineUser(data, up.C)
	}
}

func (this *SmsProcess) SendMesToEachOnlineUser(data []byte, c net.Conn) {
	tf := &utils.Transfer{
		C: c,
	}
	err := tf.WritePkg(data)
	if err != nil {
		fmt.Println("转发消息失败:", err)
	}
}
