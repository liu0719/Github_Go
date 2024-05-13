package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/", func(ctx *gin.Context) {
		ctx.String(200, "值:%v", "你好gin")
	})
	r.Run("10.50.142.216:8888")
}
