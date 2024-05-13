package model

type CartItem struct {
	CartItemId int64   //购物项的ID
	Book       *Book   //购物项的图书信息
	Count      int64   //购物项的数量
	Amount     float64 //购物项的金额小计，通过计算得到
	CartId     string  //当钱购物项属于哪个购物车
}

//获取购物项的金额小计
func (cartitem *CartItem) GetAmount() float64 {
	price := cartitem.Book.Price
	return float64(cartitem.Count) * price
}
