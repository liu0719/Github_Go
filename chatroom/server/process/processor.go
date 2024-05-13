package process

import (
	"chatroom/client/utils"
	"chatroom/common/message"
	"fmt"
	"io"
	"net"
)

type Processor struct {
	C net.Conn
}

func (this *Processor) Process() (err error) {
	for {
		tf := &utils.Transfer{
			C: this.C,
		}
		mes, err := tf.ReadPkg()
		if err != nil {
			if err == io.EOF {
				fmt.Println(this.C.RemoteAddr().String(), ":", err)
				return err
			}
			fmt.Println(err)
			return err
		}
		err = this.serverProcessMes(&mes)
	}
}

func (this *Processor) serverProcessMes(mes *message.Message) (err error) {
	switch mes.Type {
	case message.LoginMesType:
		//登录信息处理
		up := &UserProcess{
			C: this.C,
		}
		err = up.serverProcessLogin(mes)
		if err != nil {
			fmt.Println(err)
			return
		}
	case message.RegisterMesType:
		//注册信息处理
		up := &UserProcess{
			C: this.C,
		}
		err = up.ServerProcessRegister(mes)
		if err != nil {
			fmt.Println(err)
			return
		}
	case message.SmsMesType:
		//发送消息处理
		//创建一个SmsProcess实例完成转发世界消息
		sm := &SmsProcess{}
		sm.SendGroupMes(mes)
	case message.NotifyUserStatusMesType:
		//处理离线消息
		up := &UserProcess{
			C: this.C,
		}
		up.ServerProcessOffline(mes)
	case message.PersonSmsMesType:
		//私信消息处理
		fmt.Println(mes.Data)
	default:
		fmt.Println(this.C.RemoteAddr().String(), ":发送的消息类型不存在,无法处理")
	}
	return
}
