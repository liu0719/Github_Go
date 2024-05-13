package routers

import (
	"bookstore/controllers"

	"github.com/gin-gonic/gin"
)

func AddAdminRouter(r *gin.Engine) {
	r.Static("/admin/static/", "./views/static/")
	adminRouter := r.Group("/admin")
	{
		adminRouter.GET("/", controllers.AdminHandler{}.Admin)
		adminRouter.GET("/books", controllers.BookHandler{}.GetPageBooks)
		adminRouter.GET("/toUpdatebookPage", controllers.BookHandler{}.UpDateOrAddBookPage)
		adminRouter.POST("/UpdateOrAddBook", controllers.BookHandler{}.UpDateOrAddBook)
		adminRouter.GET("/deleteBook", controllers.BookHandler{}.DeleteBook)
		adminRouter.GET("/getOrders", controllers.OrderHandler{}.GetOrders)
		adminRouter.GET("/getOrderInfo", controllers.AdminHandler{}.GetOrderInfo)
		adminRouter.GET("/sendOrder", controllers.OrderHandler{}.SendOrder)

	}
}
