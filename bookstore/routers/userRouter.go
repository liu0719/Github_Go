package routers

import (
	"bookstore/controllers"

	"github.com/gin-gonic/gin"
)

func AddUserRouter(r *gin.Engine) {
	r.GET("/", controllers.UserHandler{}.Index)
	//主页图书价格和分页查询
	r.GET("/getPageBooksByPrice", controllers.BookHandler{}.GetPageBooksByPrice)
	UserRouter := r.Group("/user")
	{
		UserRouter.GET("/login", controllers.UserHandler{}.Login)
		UserRouter.GET("/regist", controllers.UserHandler{}.Regist)
	}
}
