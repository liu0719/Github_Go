package dao

import (
	"go_web/bookstore/model"
	"go_web/bookstore/utils"
)

// AddCartItem 向购物项列表中插入购物项
func AddCartItem(cartItem *model.CartItem) error {
	sql := "insert into cart_items(count,amount,book_id,cart_id)values(?,?,?,?)"
	_, err := utils.Db.Exec(sql, cartItem.Count, cartItem.GetAmount(), cartItem.Book.Id, cartItem.CartId)
	if err != nil {
		return err
	}
	return nil
}

// GetCartItemByBookId根据图书的id获取对应的购物项
func GetCartItemByBookIdAndCartId(bookId string, cartId string) (*model.CartItem, error) {
	// 写sql
	sql := "select id,count,amount,cart_id from cart_items where book_id = ? and cart_id = ?"
	row := utils.Db.QueryRow(sql, bookId, cartId)
	cartItem := &model.CartItem{}
	err := row.Scan(&cartItem.CartItemId, &cartItem.Count, &cartItem.Amount, &cartItem.CartId)
	if err != nil {
		return nil, err
	}
	cartItem.Book, err = GetBookById(bookId)
	if err != nil {
		return nil, err
	}
	return cartItem, nil
}

// UpDateBookCount 根据图书的ID以及图书的数量更新购物车中的数量和金额小计
func UpDateBookCount(cartItem *model.CartItem) error {
	sql := "update cart_items set count = ?, amount = ? where book_id = ? and cart_id = ?"
	_, err := utils.Db.Exec(sql, cartItem.Count, cartItem.GetAmount(), cartItem.Book.Id, cartItem.CartId)
	if err != nil {
		return err
	}
	return nil
}

// GetCartItemByCartId 根据购物车的Id获取图书中的所有购物项
func GetCartItemByCartId(cartId string) (cartItemSlice []*model.CartItem, err error) {
	sql := "select id,count,amount,book_id,cart_id from cart_items where cart_id = ?"
	rows, err := utils.Db.Query(sql, cartId)
	if err != nil {
		return
	}
	//设置一个变量，接收bookId
	var bookId string
	for rows.Next() {
		cartItem := &model.CartItem{}
		err = rows.Scan(&cartItem.CartItemId, &cartItem.Count, &cartItem.Amount, &bookId, &cartItem.CartId)
		if err != nil {
			return nil, err
		}
		cartItem.Book, err = GetBookById(bookId)
		if err != nil {
			return nil, err
		}

		cartItemSlice = append(cartItemSlice, cartItem)
	}
	return
}

// 删除单个购物项
func DeleteCartItemById(cartItemId string) error {
	sql := "delete from cart_items where id = ?"
	_, err := utils.Db.Exec(sql, cartItemId)
	if err != nil {
		return err
	}
	return nil
}

// DeleteCartItemsByCartId 根据购物车的ID删除所有的购物项
func DeleteCartItemsByCartId(cartId string) error {
	sql := "delete from cart_items where cart_id = ?"
	_, err := utils.Db.Exec(sql, cartId)
	if err != nil {
		return err
	}
	return nil
}

// DeleteCartById 根据购物车ID删除购物车
func DeleteCartById(cartId string) error {
	//删除购物车之前需要删除所有的购物享
	err := DeleteCartItemsByCartId(cartId)
	if err != nil {
		return err
	}
	sql := "delete from carts where id = ?"
	_, err = utils.Db.Exec(sql, cartId)
	if err != nil {
		return err
	}
	return nil
}
