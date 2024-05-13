package routers

import (
	"go_gin/server_02/controllers/admin"
	"go_gin/server_02/middlewares"

	"github.com/gin-gonic/gin"
)

func AddadminRouters(r *gin.Engine) {
	adminRouter := r.Group("/admin", middlewares.InitIndex)
	// 也可以这样配置
	// adminRouter.Use(middlewares.InitIndex)
	{
		//一个请求可以传递多个处理函数，配置路由前后都可以，称为中间件，就像vue的生命周期钩子一个道理
		adminRouter.GET("/", admin.IndexController{}.Index)

		adminRouter.GET("/users", admin.UserController{}.Index)
		adminRouter.GET("/users/add", admin.UserController{}.Add)
		adminRouter.GET("/users/addSame", admin.UserController{}.AddSame)
		adminRouter.GET("/users/edit", admin.UserController{}.Edit)
		adminRouter.GET("/users/delete", admin.UserController{}.Delete)
		adminRouter.POST("/users/doUpload", admin.UserController{}.DoUpload)
		adminRouter.POST("/users/doUploadSame", admin.UserController{}.DoUploadSame)
		adminRouter.POST("/users/doEdit", admin.UserController{}.DoEdit)

		adminRouter.GET("/artcile", admin.ArtcileController{}.Index)
		adminRouter.GET("/artcile/add", admin.ArtcileController{}.Add)
		adminRouter.GET("/artcile/edit", admin.ArtcileController{}.Edit)
	}
}
