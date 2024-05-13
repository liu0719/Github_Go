package routers

import (
	"go_gin/server_02/controllers/itying"

	"github.com/gin-gonic/gin"
)

func AddDefaultRouters(r *gin.Engine) {
	defaultRouters := r.Group("/")
	{
		defaultRouters.GET("/", itying.DefaultController{}.Index)
		defaultRouters.GET("/news", itying.DefaultController{}.News)
		defaultRouters.GET("/shop", itying.DefaultController{}.Shop)
		defaultRouters.GET("/delete", itying.DefaultController{}.Delete)
	}
}
