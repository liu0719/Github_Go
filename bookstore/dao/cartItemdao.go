package dao

import (
	"bookstore/models"
	"bookstore/utils"
)

// GetCartItemByBookId根据图书的id获取对应的购物项
func GetCartItemByBookIdAndCartId(bookId string, cartId string) (*models.CartItem, error) {
	cartItem := &models.CartItem{}
	result := utils.DB.Where("book_id = ? and cart_id= ?", bookId, cartId).Preload("Book").First(cartItem)
	if result.Error != nil {
		return nil, result.Error
	}
	return cartItem, nil
}

// 根据图书的ID以及图书的数量更新购物车中的数量和金额小计
func UpDateBookCount(cartItem *models.CartItem) error {
	result := utils.DB.Where("book_id = ? and cart_id = ?", cartItem.BookId, cartItem.CartId).Updates(&models.CartItem{Count: cartItem.Count, Amount: cartItem.GetAmount()})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// 向购物项列表中插入购物项
func AddCartItem(cartItem *models.CartItem) error {
	result := utils.DB.Create(&models.CartItem{Count: cartItem.Count, Amount: cartItem.GetAmount(), BookId: cartItem.BookId, CartId: cartItem.CartId})
	if result.Error != nil {

		return result.Error
	}
	return nil
}

// 根据购物车ID删除购物车
func DeleteCartById(cartId string) error {
	//删除购物车之前需要删除所有的购物享
	err := DeleteCartItemsByCartId(cartId)
	if err != nil {
		return err
	}
	result := utils.DB.Where("id= ?", cartId).Delete(&models.Cart{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// DeleteCartItemsByCartId 根据购物车的ID删除所有的购物项
func DeleteCartItemsByCartId(cartId string) error {
	result := utils.DB.Where("cart_id = ?", cartId).Delete(&models.CartItem{})

	if result.Error != nil {
		return result.Error
	}
	return nil
}

// 删除单个购物项
func DeleteCartItemById(cartItemId string) error {
	result := utils.DB.Where("id=?", cartItemId).Delete(&models.CartItem{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}
