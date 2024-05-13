package dao

import (
	"go_web/bookstore/model"
	"go_web/bookstore/utils"
)

// 向购物车表中插入购物车
func AddCart(cart *model.Cart) error {
	//cartItem的值为空时，打开一下循环
	// for _, v := range cart.CartItems {
	// 	v.Amount = v.GetAmount()
	// }
	sql := "insert into carts(id,total_count,total_amount,user_id) values(?,?,?,?)"
	_, err := utils.Db.Exec(sql, cart.CartId, cart.GetTotalCount(), cart.GetTotlalAmount(), cart.UserId)
	if err != nil {
		return err
	}
	//获取购物车中的所有购物项
	cartItems := cart.CartItems
	for _, cartItem := range cartItems {
		//建构无想插入到数据库中
		AddCartItem(cartItem)
	}
	return nil
}

// GeyCartByUserId根据用户ID从数据库中查询对应的购物车
func GetCartByUserId(userId int) (*model.Cart, error) {
	sql := "select id,total_count,total_amount,user_id from carts where user_id = ?"
	row := utils.Db.QueryRow(sql, userId)
	cart := &model.Cart{}
	err := row.Scan(&cart.CartId, &cart.TotalCount, &cart.TotalAmount, &cart.UserId)
	if err != nil {
		return nil, err
	}

	//获取当前购物车的所有购物项
	cartItemSlice, err := GetCartItemByCartId(cart.CartId)
	if err != nil {
		return nil, err
	}
	cart.CartItems = cartItemSlice
	return cart, nil
}

// UpDateCart 更新购物车中的图书的总数量和总金额
func UpDateCart(cart *model.Cart) error {
	sql := "update carts set total_count  = ?,total_amount = ? where id = ?"
	_, err := utils.Db.Exec(sql, cart.GetTotalCount(), cart.GetTotlalAmount(), cart.CartId)
	if err != nil {
		return err
	}
	return nil
}
