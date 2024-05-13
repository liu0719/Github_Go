package message

const (
	LoginMesType            = "LoginMes"
	LoginResMesType         = "LoginResMes"
	RegisterMesType         = "RegisterMes"
	RegisterResMesType      = "RegisterResMes"
	NotifyUserStatusMesType = "NotifyUserStatusMes"
	SmsMesType              = "SmsMes"
)

//定义几个长亮表示用户状态
const (
	UserOnline = iota
	UserOffline
	UserBusyStatus
)

//通用消息类型
type Message struct {
	Type string `json:"type"` //消息类型
	Data string `json:"data"` //消息内容
}

// 客户端发送消息
type LoginMes struct {
	UserId   int    `json:"userId"`   //用户账号
	UserPwd  string `json:"userPwd"`  //用户密码
	UserName string `json:"userName"` //用户名
}

// 服务端返回消息
type LoginResMes struct {
	Code     int    `json:"type"`  //返回状态码
	Error    string `json:"error"` //返回错误信息
	UserName string `json:"username"`
	UsersId  []int  //增加字段，保存用户id的切片
}

type RegisterMes struct {
	User User `json:"user"`
}

type RegisterResMes struct {
	Code  int    `json:"type"`  //返回状态码表示该用户已经占用，200表示注册成功
	Error string `json:"error"` //返回错误信息
}

//为了配合服务器推送通知，用户状态变化的消息
type NotifyUserStatusMes struct {
	UserId int `json:"userId"` //用户Id
	Status int `json:"status"` //用户状态
}

//增加一个SmsMes//发送消息
type SmsMes struct {
	Content string `json:"content"`
	User           //匿名结构体，继承
}

// SmsReMes
