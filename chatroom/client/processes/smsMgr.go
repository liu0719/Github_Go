package processes

import (
	"chatroom/common/message"
	"encoding/json"
	"fmt"
)

func outputGroup(mes *message.Message) { //确定类型
	//显示即可

	// 1.反序列化
	var smsMes message.SmsMes
	err := json.Unmarshal([]byte(mes.Data), &smsMes)
	if err != nil {
		fmt.Println("fan序列化失败:", err.Error())
		return
	}

	//显示
	content := fmt.Sprintf("%v(ID:%v):%v", smsMes.UserName, smsMes.UserId, smsMes.Content)
	fmt.Println(content)
	fmt.Println()
}
