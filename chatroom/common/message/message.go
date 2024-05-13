package message

var (
	LoginMesType            = "LoginMes"
	LoginResMesType         = "LoginResMes"
	RegisterMesType         = "RegisterMes"
	RegisterResMesType      = "RegisterResMes"
	NotifyUserStatusMesType = "NotifyUserStatusMes"
	SmsMesType              = "SmsMes"
	PersonSmsMesType        = "PersonSmsMes"
	SendNewMapType          = "SendNewMap"
)

//定义几个长亮表示用户状态
const (
	UserOnline = iota
	UserOffline
	UserBusyStatus
)

type Message struct {
	Type string `json:"type"` //类型头
	Data string `json:"data"` //消息体
}

type LoginMes struct {
	User
}

type LoginResMes struct {
	Code  int    `json:"code"`  //返回状态码
	Error string `json:"error"` //返回的错无
	User  `json:"user"`
	Users []*User // 增加字段，保存用户的切片
}
type RegisterMes struct {
	User `json:"user"`
}

type RegisterResMes struct {
	Code  int    `json:"type"`  //返回状态码表示该用户已经占用，200表示注册成功
	Error string `json:"error"` //返回错误信息
}

//为了配合服务器推送通知，用户状态变化的消息
type NotifyUserStatusMes struct {
	UserId   int    `json:"userId"` //用户Id
	Status   int    `json:"status"` //用户状态
	UserName string `json:"userName"`
}

//增加一个SmsMes//发送消息
type SmsMes struct {
	Content string `json:"content"`
	User           //匿名结构体，继承
}

type PersonSmsMes struct {
	Content  string `json:"content"`
	User            //匿名结构体，继承
	ToUserId int    //发给谁
}

//服务器发送给客户端的列表，使其更新
type OnlineUsers map[int]*User
