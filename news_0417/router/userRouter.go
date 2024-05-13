package router

import (
	"news_0417/handler"

	"github.com/gin-gonic/gin"
)

func UserRouter(r *gin.Engine) {
	userRouter := r.Group("/user")
	userRouter.GET("/addNewsPage", handler.UserHandler{}.AddNewsPage)
	userRouter.POST("/addNews", handler.UserHandler{}.AddNews)
	userRouter.GET("/delete", handler.UserHandler{}.DeleteNew)
	userRouter.GET("/update", handler.UserHandler{}.UpdateNew)
	userRouter.POST("/selectAuthor", handler.UserHandler{}.SelectNewsByAuhtor)
	userRouter.POST("/selectClass", handler.UserHandler{}.SelectNewsByClass)

}
