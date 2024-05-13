package model

import (
	message "go_network/chatroom/common/messgae"
	"net"
)

//因为在客户端，我们很多地方会使用到curUser,将其作为一个全局的
type CurUser struct {
	C net.Conn
	message.User
}
