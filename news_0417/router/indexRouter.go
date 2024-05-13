package router

import (
	"news_0417/handler"

	"github.com/gin-gonic/gin"
)

func IndexRouter(r *gin.Engine) {
	userRouter := r.Group("/")
	{
		userRouter.GET("/", handler.UserHandler{}.Index)
	}

}
