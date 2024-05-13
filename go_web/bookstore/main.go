package main

import (
	"go_web/bookstore/controller"
	"net/http"
)

func main() {
	//设置处理静态资源
	//可以直达html页面,不需要通过引擎
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("bookstore/views/static/"))))
	//直接去html页面
	http.Handle("/pages/", http.StripPrefix("/pages/", http.FileServer(http.Dir("bookstore/views/pages/"))))

	http.HandleFunc("/", controller.IndexHandler)

	//去登录
	http.HandleFunc("/login", controller.Login)

	//登出
	http.HandleFunc("/logout", controller.Logout)

	//去注册
	http.HandleFunc("/register", controller.Register)

	// 验证Ajax用户名是否可用
	http.HandleFunc("/checkUserName", controller.CheckUserName)

	//获取所有图书
	// http.HandleFunc("/getBooks", controller.GetBooks)

	//获取带分页的图书信息
	http.HandleFunc("/getPageBooks", controller.GetPageBooks)

	//根据价格获取书
	http.HandleFunc("/getPageBooksByPrice", controller.GetPageBooksByPrice)

	//添加图书
	// http.HandleFunc("/addBook", controller.AddBook)

	//删除图书
	http.HandleFunc("/deleteBook", controller.DeleteBook)

	//修改图书
	http.HandleFunc("/toUpdateBookPage", controller.ToUpdateBookPage)

	// 更新图书
	http.HandleFunc("/UpdateOrAddBook", controller.UpdateOrAddBook)

	//添加图书到购物车
	http.HandleFunc("/addBook2Cart", controller.AddBook2Cart)

	//删除购物车（清空）
	http.HandleFunc("/deleteCart", controller.DeleteCart)

	//删除购物项
	http.HandleFunc("/deleteCartItem", controller.DeleteCartItem)

	//获取购物车
	http.HandleFunc("/getCartInfo", controller.GetCartInfo)

	//更新购物车
	http.HandleFunc("/updateCartItem", controller.UpdateCartItem)

	//去结账
	http.HandleFunc("/checkout", controller.CheckOut)

	//获取所有订单
	http.HandleFunc("/getOrders", controller.GetOrders)

	//获取订单项请，即订单所对应的所有订单项
	http.HandleFunc("/getOrderInfo", controller.GetOrderInfo)

	//获取我的订单
	http.HandleFunc("/getMyOrder", controller.GetMyOrders)

	// 发货
	http.HandleFunc("/sendOrder", controller.SendOrder)

	//确认收货
	http.HandleFunc("/takeOrder", controller.TakeOrder)

	//图标
	http.HandleFunc("/favicon.ico", controller.Favicon)

	// server := http.Server{
	// 	Addr:        "10.50.142.216:8888",
	// 	Handler:     nil,
	// 	ReadTimeout: 10 * time.Second,
	// }
	// server.ListenAndServe()
	http.ListenAndServe(":8888", nil)
}
