package models

// OrderItem 结构体
type OrderItem struct {
	Id      int     //订单项的Id
	Count   int     //订单项中图书的数量
	Amount  float64 //订单的金额小计
	Title   string
	Author  string
	Price   float64
	ImgPath string
	OrderId string
}

func (OrderItem) TableName() string {
	return "order_items"
}
