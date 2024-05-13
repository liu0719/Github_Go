package model

// Session结构体
type Session struct {
	SessionId string
	UserName  string
	UserId    int
	Cart      *Cart
	OrderId   string
	Orders    []*Order
}
