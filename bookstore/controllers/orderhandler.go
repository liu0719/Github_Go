package controllers

import (
	"bookstore/dao"
	"bookstore/models"
	"bookstore/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
}

func (OrderHandler) CheckOut(c *gin.Context) {
	//获取session
	_, session := dao.IsLogin(c)
	//获取用户ID
	userId := session.UserId
	//根据ID获取购物车
	cart, err := dao.GetCartByUserId(userId)
	if err != nil {
		return
	}
	//调用utils包，获取当前时间
	nowtime := time.Now().Format("2006-01-02 15:04:05")
	//生成订单号
	orderId := utils.CreateUUID()
	// 创建Order
	order := &models.Order{
		Id:          orderId,
		CreateTime:  nowtime,
		TotalCount:  cart.TotalCount,
		TotalAmount: cart.TotalAmount,
		State:       0,
		UserId:      userId,
	}

	//将订单保存到数据库
	dao.AddOrder(order)
	//保存订单项
	for _, v := range cart.CartItems {
		//获取购物车中的所有购物享
		orderItem := &models.OrderItem{
			Count:   v.Count,
			Amount:  v.Amount,
			Title:   v.Book.Title,
			Author:  v.Book.Author,
			Price:   v.Book.Price,
			ImgPath: v.Book.ImgPath,
			OrderId: orderId,
		}
		//将购物项保存到数据库
		dao.AddOrderItem(orderItem)
		//更新当前购物项总图书的库存和销量
		book := v.Book
		book.Sales = book.Sales + int(v.Count)
		book.Stock = book.Stock - int(v.Count)
		// 跟新图树的库存和销量
		dao.UpdateBook(book)
	}
	//生成订单后删除购物车
	dao.DeleteCartById(cart.Id)
	//将订单号设置到session哩
	session.OrderId = orderId
	//解析模板去订单页面
	c.HTML(http.StatusOK, "checkout.html", session)
}

// 用户获取属于自己ID的订单
func (OrderHandler) GetMyOrders(c *gin.Context) {
	_, session := dao.IsLogin(c)
	userId := session.UserId
	//调用dao中获取所有订单的函数
	orders, _ := dao.GetOrdersByUserId(userId)

	session.Orders = orders
	//解析模板
	c.HTML(http.StatusOK, "order.html", session)
}

// getOrderInfo 获取订单详情
func (OrderHandler) GetOrderInfo(c *gin.Context) {
	flag, session := dao.IsLogin(c)
	if !flag {
		c.Redirect(http.StatusFound, "/user/login")
	}
	// 获取订单号
	orderId := c.Query("orderId")
	orders, err := dao.GetOrderItemByOrderId(orderId)
	if err != nil {
		return
	}
	c.HTML(http.StatusOK, "order_info.html", gin.H{
		"orders":   orders,
		"isAdmin":  false,
		"Username": session.Username,
	})
}
func (OrderHandler) GetOrders(c *gin.Context) {
	//调用dao中获取所有订单的函数
	orders, _ := dao.GetOrders()
	c.HTML(http.StatusOK, "order_manager.html", orders)
}

// 发货
func (OrderHandler) SendOrder(c *gin.Context) {
	//获取要发货的订单
	orderId := c.Query("orderId")
	err := dao.UpdateOrderState(orderId, 1)
	if err != nil {
		return
	}
	//调用getOrders函数再次查询所有订单
	OrderHandler{}.GetOrders(c)
}

// 确认收货
func (OrderHandler) TakeOrder(c *gin.Context) {
	//获取要发货的订单
	orderId := c.Query("orderId")
	err := dao.UpdateOrderState(orderId, 2)
	if err != nil {
		return
	}
	//调用getMyOrders函数再次查询我的订单
	OrderHandler{}.GetMyOrders(c)

}
