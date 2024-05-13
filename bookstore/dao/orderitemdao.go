package dao

import (
	"bookstore/models"
	"bookstore/utils"
	"errors"
)

// AddOrderItem 项数据库插入订单项
func AddOrderItem(orderItem *models.OrderItem) error {
	result := utils.DB.Create(orderItem)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// 根据订单号获取所有的订单项
func GetOrderItemByOrderId(orderId string) ([]models.OrderItem, error) {
	orderItems := []models.OrderItem{}
	result := utils.DB.Where("order_id =?", orderId).Find(&orderItems)
	if len(orderItems) == 0 {
		return nil, errors.New("没有订单项")
	}
	if result.Error != nil {
		return nil, result.Error
	}

	return orderItems, nil
}
