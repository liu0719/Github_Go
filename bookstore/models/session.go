package models

// Session结构体
type Session struct {
	SessionId string
	Username  string
	UserId    int
	Cart      *Cart
	OrderId   string
	Orders    []Order
}

//向数据库临时查询的Session
type SessionQuery struct {
	SessionId string
	Username  string
	UserId    int
}

func (SessionQuery) TableName() string {
	return "sessions"
}
