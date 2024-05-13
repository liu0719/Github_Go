package models

type CartItem struct {
	Id     int     //购物项的ID
	Book   *Book   `gorm:"foreignKey:BookId"` //购物项的图书信息
	Count  int     //购物项的数量
	Amount float64 //购物项的金额小计，通过计算得到
	CartId string  //当钱购物项属于哪个购物车
	BookId int     //购物相中书籍的ID
}

func (CartItem) TableName() string {
	return "cart_items"
}

// 获取购物项的金额小计
func (cartitem *CartItem) GetAmount() float64 {
	price := cartitem.Book.Price
	return float64(cartitem.Count) * price
}
