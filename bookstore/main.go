package main

import (
	"bookstore/routers"

	"github.com/gin-gonic/gin"
	"github.com/thinkerou/favicon"
)

func main() {
	//开启默认路由
	r := gin.Default()
	//favicon图标
	r.Use(favicon.New("views/favicon.ico"))
	//配置模板路径
	r.LoadHTMLGlob("views/pages/**/*")
	// 静态资源配置
	r.Static("/static/", "./views/static/")
	//访客路由分组
	routers.AddUserRouter(r)
	// 管理员路由分组
	routers.AddAdminRouter(r)
	//api接口路由分组
	routers.AddApiRouter(r)
	//启动监听
	r.Run(":80")
}
