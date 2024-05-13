package controller

import (
	"go_web/bookstore/dao"
	"go_web/bookstore/model"
	"go_web/bookstore/utils"
	"net/http"
	"text/template"
	"time"
)

// 去结账
func CheckOut(w http.ResponseWriter, r *http.Request) {
	//获取session
	_, session := dao.IsLogin(r)
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
	order := &model.Order{
		OrderId:     orderId,
		CreateTime:  nowtime,
		TotalCount:  cart.TotalCount,
		TotalAmount: cart.TotalAmount,
		State:       0,
		UserId:      int64(userId),
	}

	//将订单保存到数据库
	dao.AddOrder(order)
	//保存订单项
	for _, v := range cart.CartItems {
		//获取购物车中的所有购物享
		orderItem := &model.OredrItem{
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
	dao.DeleteCartById(cart.CartId)
	//将订单号设置到session哩
	session.OrderId = orderId
	//解析模板去订单页面
	t := template.Must(template.ParseFiles("bookstore/views/pages/cart/checkout.html"))
	t.Execute(w, session)
}

func GetOrders(w http.ResponseWriter, r *http.Request) {
	//调用dao中获取所有订单的函数
	orders, _ := dao.GetOrders()
	//解析模板
	t := template.Must(template.ParseFiles("bookstore/views/pages/order/order_manager.html"))
	t.Execute(w, orders)
}

func GetMyOrders(w http.ResponseWriter, r *http.Request) {
	_, session := dao.IsLogin(r)
	userId := session.UserId
	//调用dao中获取所有订单的函数
	orders, _ := dao.GetOrdersByUserId(userId)
	session.Orders = orders
	//解析模板
	t := template.Must(template.ParseFiles("bookstore/views/pages/order/order.html"))
	t.Execute(w, session)
}

// getOrderInfo 获取订单详情
func GetOrderInfo(w http.ResponseWriter, r *http.Request) {
	// 获取订单号
	orderId := r.FormValue("orderId")
	orders, err := dao.GetOrderItemByOrderId(orderId)
	if err != nil {
		return
	}
	t := template.Must(template.ParseFiles("bookstore/views/pages/order/order_info.html"))
	t.Execute(w, orders)
}

// 发货
func SendOrder(w http.ResponseWriter, r *http.Request) {
	//获取要发货的订单
	orderId := r.FormValue("orderId")
	err := dao.UpdateOrderState(orderId, 1)
	if err != nil {
		return
	}
	//调用getOrders函数再次查询所有订单
	GetOrders(w, r)
}

// 确认收货
func TakeOrder(w http.ResponseWriter, r *http.Request) {
	//获取要发货的订单
	orderId := r.FormValue("orderId")
	err := dao.UpdateOrderState(orderId, 2)
	if err != nil {
		return
	}
	//调用getMyOrders函数再次查询我的订单
	GetMyOrders(w, r)
}
