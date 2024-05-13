package message

//定义用户的结构体
type User struct {
	//确定字段
	// 用户信息的json字符串和结构体的字段对应的tag名字要一致
	UserId     int    `json:"userId"`
	UserPwd    string `json:"userPwd"`
	UserName   string `json:"userName"`
	UserStatus int    `json:"userStatus"` //用户状态
	Sex        string `json:"sex"`
}
