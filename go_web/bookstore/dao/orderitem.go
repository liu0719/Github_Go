package dao

import (
	"go_web/bookstore/model"
	"go_web/bookstore/utils"
)

// AddOrderItem 项数据库插入订单项
func AddOrderItem(orderItem *model.OredrItem) error {
	sql := "insert into order_items(count,amount,title,author,price,img_path,order_id)values(?,?,?,?,?,?,?)"
	_, err := utils.Db.Exec(sql, orderItem.Count, orderItem.Amount, orderItem.Title, orderItem.Author, orderItem.Price, orderItem.ImgPath, orderItem.OrderId)
	if err != nil {
		return err
	}
	return nil
}

// 根据订单号获取所有的订单项
func GetOrderItemByOrderId(orderId string) ([]*model.OredrItem, error) {
	sql := "select id,count,amount,title,author,price,img_path,order_id from order_items where order_id = ?"
	rows, err := utils.Db.Query(sql, orderId)
	if err != nil {
		return nil, err
	}
	var orderItemSlice []*model.OredrItem
	for rows.Next() {
		orderItem := &model.OredrItem{}
		rows.Scan(&orderItem.OrderItemId, &orderItem.Count, &orderItem.Amount, &orderItem.Title, &orderItem.Author, &orderItem.Price, &orderItem.ImgPath, &orderItem.OrderId)
		orderItemSlice = append(orderItemSlice, orderItem)
	}
	return orderItemSlice, nil
}
