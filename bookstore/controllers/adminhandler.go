package controllers

import (
	"bookstore/dao"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AdminHandler struct {
}

func (AdminHandler) Admin(c *gin.Context) {

	c.HTML(http.StatusOK, "manager.html", "")
}

// getOrderInfo 获取订单详情
func (AdminHandler) GetOrderInfo(c *gin.Context) {
	// 获取订单号
	orderId := c.Query("orderId")
	orders, err := dao.GetOrderItemByOrderId(orderId)
	if err != nil {
		return
	}
	c.HTML(http.StatusOK, "order_info.html", gin.H{
		"orders":  orders,
		"isAdmin": true,
	})
}
