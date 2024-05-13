package processes

import (
	"encoding/json"
	"fmt"
	"go_network/chatroom/client/utils"
	message "go_network/chatroom/common/messgae"
)

type SmsProcess struct {
}

//发送群聊的消息
func (this *SmsProcess) SendGroup(content string, userName string) (err error) {
	//1.创建发送的Mes
	var mes message.Message
	mes.Type = message.SmsMesType

	//2.创建SmsMes
	var smsMes message.SmsMes
	smsMes.Content = content               //内容
	smsMes.UserId = CurUser.UserId         //ID
	smsMes.UserStatus = CurUser.UserStatus //状态
	smsMes.UserName = userName
	//3.序列化SmsMes
	data, err := json.Marshal(smsMes)
	if err != nil {
		fmt.Println("序列化失败:", err.Error())
		return
	}

	//4.将序列化后的SmsMes(data),赋给Mes.Data
	mes.Data = string(data)

	//5.将mes序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("序列化失败:", err.Error())
		return
	}

	//5.发送序列化后的mes给服务器
	tf := &utils.Transfer{
		C: CurUser.C,
	}
	err = tf.WriterPkg(data)
	if err != nil {
		fmt.Println("发送失败:", err.Error())
		return
	}
	return
}