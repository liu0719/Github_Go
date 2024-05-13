package routers

import "github.com/gin-gonic/gin"

func AddapiRouters(r *gin.Engine) {
	apiRouters := r.Group("/api")
	{
		apiRouters.GET("/", func(c *gin.Context) {
			c.String(200, "api接口")
		})
		apiRouters.GET("/userlist", func(c *gin.Context) {
			c.String(200, "api接口---userlist")
		})
		apiRouters.GET("/plist", func(c *gin.Context) {
			c.String(200, "api接口----plist")
		})
	}
}
