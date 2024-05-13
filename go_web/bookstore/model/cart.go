package model

type Cart struct {
	CartId      string      //购物车的ID
	CartItems   []*CartItem //购物车中所有的购物项
	TotalCount  int64       //购物项总计
	TotalAmount float64     //购物钱数总计
	UserId      int         //当前购物车所属的用户
	UserName    string      //返回时显示用户名
}

// GetTotalCount 获取购物车中的图书的数量
func (c *Cart) GetTotalCount() int64 {
	var totalCount int64
	//便利购物车中所有的购物想切片
	for _, v := range c.CartItems {
		totalCount += v.Count
	}
	return totalCount
}

// GetTotalAmount  获取购物车中的图书的总金额
func (c Cart) GetTotlalAmount() float64 {
	var totalAmount float64
	//便利购物车中所有的购物想切片
	for _, v := range c.CartItems {
		totalAmount += v.GetAmount()
	}
	return totalAmount
}
