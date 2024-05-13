package main

import (
	"text/template"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/thinkerou/favicon"
)

// 时间戳转换成日期
func UnixToTime(timestamp int64) string {
	t := time.Unix(timestamp, 0)
	return t.Format("2006-01-02 15:04:05")
}
func Println(str1 string, str2 string) string {
	return str1 + str2
}
func main() {
	r := gin.Default()

	//自定义模板函数，一定要放在加载模板前
	r.SetFuncMap(template.FuncMap{
		"UnixToTime": UnixToTime,
		"Println":    Println,
	})

	//配置模板路径
	r.LoadHTMLFiles("static/music/index.html")
	//静态资源
	r.Static("/static/", "./static/")

	//图标
	r.Use(favicon.New("favicon.ico"))

	r.GET("/", func(c *gin.Context) {
		//返回jsonp,区别就是可以传入回调函数
		// jsonp主要用来解决跨域问题
		// a := &Article{
		// 	Title:   "三国演义-jsonp",
		// 	Desc:    "我是描述",
		// 	Content: "东汉末年",
		// }
		c.HTML(200, "index.html", gin.H{
			"msg":   "我是后台数据",
			"title": "我是标题-->",
			"score": 78,
			"date":  time.Now().Unix(),
		})
	})

	//监听端口，启动
	r.Run("10.50.142.216:8888")
}
