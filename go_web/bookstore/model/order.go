package model

//Order结构体
type Order struct {
	OrderId     string  //订单号
	CreateTime  string  //生成时间
	TotalCount  int64   //订单总数量
	TotalAmount float64 //订单中的总金额
	State       int64   //订单状态 0 未发货  1 已发货 2 交易完成
	UserId      int64   //订单所属的用户

}

//发货
func (o *Order) NoSend() bool {
	return o.State == 0
}

// 收货
func (o *Order) SendComplate() bool {
	return o.State == 1
}

//交易完成
func (o *Order) Complate() bool {
	return o.State == 2
}
