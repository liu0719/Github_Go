package dao

import (
	"bookstore/models"
	"bookstore/utils"
	"errors"
)

// 向数据库添加Order
func AddOrder(order *models.Order) error {
	result := utils.DB.Create(order)

	if result.Error != nil {
		return result.Error
	}
	return nil
}

// 获取数据库总所有的订单
func GetOrders() ([]models.Order, error) {
	orders := []models.Order{}
	result := utils.DB.Preload("OrderItems").Find(&orders)
	if len(orders) == 0 {
		return nil, errors.New("没有订单")
	}
	if result.Error != nil {
		return nil, result.Error
	}
	return orders, nil
}

// 根据用户ID返回该用户的订单
func GetOrdersByUserId(userId int) ([]models.Order, error) {
	orders := []models.Order{}
	result := utils.DB.Where("user_id= ?", userId).Preload("OrderItems").Find(&orders)
	if len(orders) == 0 {
		return nil, errors.New("该用户没有订单")
	}
	if result.Error != nil {
		return nil, result.Error
	}
	return orders, nil
}

// 更新订单的状态
func UpdateOrderState(orderId string, state int) error {
	result := utils.DB.Model(&models.Order{}).Where("id = ?", orderId).Update("state", state)

	if result.Error != nil {
		return result.Error
	}
	return nil
}
