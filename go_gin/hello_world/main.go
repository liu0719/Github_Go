package main

import (
	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	c.String(200, "<h1>杨佳鑫大傻b</h1>")
}
func main() {
	//1.创建路由
	r := gin.Default() //返回默认的路由引擎

	//2.绑定路由规则，执行的函数
	r.GET("/", Index)

	// 3。监听端口
	r.Run(":80")
}
