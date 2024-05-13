package main

import (
	"go_gin/server_02/models"
	"go_gin/server_02/routers"
	"text/template"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	_ "github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"github.com/thinkerou/favicon"
)

func main() {
	// 默认路由，其中会有两个中间件engine.Use(Logger(), Recovery())
	r := gin.Default()
	r.SetFuncMap(template.FuncMap{
		"UnixToTime": models.UnixToTime,
		"Println":    models.Println,
	})
	//配置session中间件
	// 创建一个基于cookie的存储引擎,secret参数是用于加密的密钥
	store := cookie.NewStore([]byte("secret"))
	//配置session的中间件 store是前面创建的存储引擎，我们可以替换成其他引擎
	r.Use(sessions.Sessions("mysession", store))

	/*
		 redis的session中间件
		// store, _ := redis.NewStore(10, "tcp", "localhost:6379", "", []byte("secret"))
		// r.Use(sessions.Sessions("mysession", store))
	*/
	r.Use(favicon.New("favicon.ico"))
	r.Static("/static/", "./static/")
	r.LoadHTMLGlob("templates/**/*")
	routers.AddDefaultRouters(r)
	routers.AddadminRouters(r)
	routers.AddapiRouters(r)

	r.Run("10.50.142.216:80")
}
