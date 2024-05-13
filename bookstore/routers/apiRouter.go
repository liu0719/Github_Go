package routers

import (
	"bookstore/controllers"

	"github.com/gin-gonic/gin"
)

func AddApiRouter(r *gin.Engine) {
	r.Static("/api/static/", "./views/static/")
	//登陆注册
	apiRouter := r.Group("/api")
	{
		//登录
		apiRouter.POST("/login", controllers.ApiHandler{}.Login)
		//注册
		apiRouter.POST("/regist", controllers.ApiHandler{}.Regist)
		//校验用户名是否可用
		apiRouter.POST("/checkUserName", controllers.ApiHandler{}.CheckUserName)
		// 登出
		apiRouter.GET("/logout", controllers.ApiHandler{}.Logout)
		// 获取购物车
		apiRouter.GET("/getCartInfo", controllers.ApiHandler{}.GetCartInfo)
		// 添加购物车
		apiRouter.POST("/addBookCart", controllers.ApiHandler{}.AddBookCart)
		// 删除购物车（清空购物车）
		apiRouter.GET("/deleteCart", controllers.ApiHandler{}.DeleteCart)
		// 删除购物项
		apiRouter.GET("/deleteCartItem", controllers.ApiHandler{}.DeleteCartItem)
		// 更新购物项
		apiRouter.POST("/updateCartItem", controllers.ApiHandler{}.UpdateCartItem)
		// 结算
		apiRouter.GET("/checkout", controllers.OrderHandler{}.CheckOut)
		//获取订单
		apiRouter.GET("/getMyOrder", controllers.OrderHandler{}.GetMyOrders)
		//获取订单详情
		apiRouter.GET("/getOrderInfo", controllers.OrderHandler{}.GetOrderInfo)
		//确认收货
		apiRouter.GET("/takeOrder", controllers.OrderHandler{}.TakeOrder)

	}

}
