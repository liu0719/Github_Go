package main

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/thinkerou/favicon"
)

type Article struct {
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

func main() {
	r := gin.Default()

	//配置模板路径
	r.LoadHTMLFiles("static/music/index.html")
	//静态资源
	r.Static("/static/", "./static/")
	r.Use(favicon.New("favicon.ico"))
	r.GET("/", func(c *gin.Context) {
		//响应一个字符串回去
		c.String(200, "值:%v。", "首页")

	})
	r.GET("/json", func(c *gin.Context) {
		//返回json
		// gin里的gin.H就是map[string]interface{}

		// c.JSON(200, map[string]interface{}{
		// 	"success": "成功",
		// })
		c.JSON(200, gin.H{
			"success": "成功",
		})
	})
	r.GET("/struct", func(c *gin.Context) {
		//返回json,可以直接返回结构体
		a := &Article{
			Title:   "三国演义",
			Desc:    "我是描述",
			Content: "东汉末年",
		}
		c.JSON(200, a)
	})
	r.GET("/jsonp", func(c *gin.Context) {
		//返回jsonp,区别就是可以传入回调函数
		// jsonp主要用来解决跨域问题
		a := &Article{
			Title:   "三国演义-jsonp",
			Desc:    "我是描述",
			Content: "东汉末年",
		}
		c.JSONP(200, a)
	})
	r.GET("/xml", func(c *gin.Context) {
		//返回jsonp,区别就是可以传入回调函数
		// jsonp主要用来解决跨域问题
		// a := &Article{
		// 	Title:   "三国演义-jsonp",
		// 	Desc:    "我是描述",
		// 	Content: "东汉末年",
		// }
		c.JSONP(200, gin.H{
			"success": "失败",
			"msg":     "gin 我是XML",
		})
	})
	r.GET("/html", func(c *gin.Context) {
		//返回jsonp,区别就是可以传入回调函数
		// jsonp主要用来解决跨域问题
		// a := &Article{
		// 	Title:   "三国演义-jsonp",
		// 	Desc:    "我是描述",
		// 	Content: "东汉末年",
		// }
		c.HTML(200, "index.html", gin.H{
			"msg":   "我是后台数据",
			"score": 78,
			"date":  time.Now().Unix(),
		})
	})
	r.Run(":8080")
}
