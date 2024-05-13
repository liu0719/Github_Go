package models

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Age      int    `json:"age"`
	Email    string `json:"email"`
	AddTime  string `json:"addTime"`
}

// 标识配置数据库的表名称
func (User) TableName() string {
	return "users"
}
