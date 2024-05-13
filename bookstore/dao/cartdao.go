package dao

import (
	"bookstore/models"
	"bookstore/utils"
	"errors"
	"strconv"
)

// 根据UserId获取购物车
func GetCartByUserId(userId int) (*models.Cart, error) {
	cart := &models.Cart{}
	result := utils.DB.Where("user_id = ?", userId).Preload("User").Preload("CartItems").First(cart)
	if result.Error != nil {
		return nil, result.Error
	}
	if cart.Id == "" {
		return nil, errors.New("该用户还没有购物车")
	}
	for _, v := range cart.CartItems {

		book, err := GetBookById(strconv.Itoa(v.BookId))
		if err != nil {
			return nil, err
		}
		v.Book = book
	}
	return cart, nil

}

// 根据购物车ID获取对应的购物项
func GetCartItemByCartId(cartId string) ([]models.CartItem, error) {
	cartItems := []models.CartItem{}
	result := utils.DB.Select("id,count,amount,book_id,cart_id").Where("cart_id=?", cartId).Preload("Book").Find(&cartItems)
	if result.Error != nil {
		return nil, result.Error
	}
	if len(cartItems) == 0 {
		return nil, errors.New("没有查找到该购物车对应的购物项")
	}
	return cartItems, nil
}

// 向购物车表中插入购物车
func AddCart(cart *models.Cart) error {
	//cartItem的值为空时，打开一下循环
	// for _, v := range cart.CartItems {
	// 	v.Amount = v.GetAmount()
	// }
	result := utils.DB.Create(&models.Cart{Id: cart.Id, TotalCount: cart.GetTotalCount(), TotalAmount: cart.TotalAmount, UserId: cart.UserId})
	if result.Error != nil {
		return result.Error
	}
	//获取购物车中的所有购物项
	cartItems := cart.CartItems
	for _, cartItem := range cartItems {
		//将购物项插入到数据库中
		AddCartItem(cartItem)
	}
	return nil
}

// UpDateCart 更新购物车中的图书的总数量和总金额
func UpDateCart(cart *models.Cart) error {
	result := utils.DB.Where("id = ?", cart.Id).Updates(&models.Cart{TotalCount: cart.GetTotalCount(), TotalAmount: cart.GetTotlalAmount()})

	if result.Error != nil {
		return result.Error
	}
	return nil
}
