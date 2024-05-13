package dao

import (
	"go_web/bookstore/model"
	"go_web/bookstore/utils"
)

func AddOrder(order *model.Order) error {
	sql := "insert into orders(id,create_time,total_count,total_amount,state,user_id)values(?,?,?,?,?,?)"
	_, err := utils.Db.Exec(sql, order.OrderId, order.CreateTime, order.TotalCount, order.TotalAmount, order.State, order.UserId)
	if err != nil {
		return err
	}
	return nil
}

// 获取数据库总所有的订单
func GetOrders() ([]*model.Order, error) {
	sql := "select id,create_time,total_count,total_amount,state,user_id from orders "

	rows, err := utils.Db.Query(sql)
	if err != nil {
		return nil, err
	}
	var orders []*model.Order
	for rows.Next() {
		order := &model.Order{}
		rows.Scan(&order.OrderId, &order.CreateTime, &order.TotalCount, &order.TotalAmount, &order.State, &order.UserId)
		orders = append(orders, order)
	}
	return orders, nil
}

// 根据用户ID返回该用户的订单
func GetOrdersByUserId(userId int) ([]*model.Order, error) {
	sql := "select id,create_time,total_count,total_amount,state,user_id from orders where user_id = ?"

	rows, err := utils.Db.Query(sql, userId)
	if err != nil {
		return nil, err
	}
	var orders []*model.Order
	for rows.Next() {
		order := &model.Order{}
		rows.Scan(&order.OrderId, &order.CreateTime, &order.TotalCount, &order.TotalAmount, &order.State, &order.UserId)
		orders = append(orders, order)
	}
	return orders, nil
}

// 更新订单的状态
func UpdateOrderState(orderId string, state int64) error {
	sql := "update orders set state = ? where id = ?"
	_, err := utils.Db.Exec(sql, state, orderId)
	if err != nil {
		return err
	}
	return nil
}
