package main

import (
	"news_0417/router"

	"github.com/gin-gonic/gin"
	"github.com/thinkerou/favicon"
)

func main() {
	r := gin.Default()
	r.Use(favicon.New("favicon.ico"))
	r.Static("/static", "./views/static/")
	r.LoadHTMLGlob("views/pages/**/*")

	//主页路由
	router.IndexRouter(r)
	//用户路由
	router.UserRouter(r)
	r.Run("localhost:80")
}
