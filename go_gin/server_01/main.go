package main

import (
	"encoding/xml"
	"net/http"
	"text/template"
	"time"

	"github.com/gin-gonic/gin"
)

type UserInfo struct {
	UserName string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
	Age      string `json:"age" form:"age"`
	Page     string `json:"page" form:"page"`
}

type Article struct {
	Title   string `json:"title" xml:"title"`
	Content string `json:"content" xml:"content"`
}

func UnixToTime(timestamp int64) string {
	t := time.Unix(timestamp, 0)
	return t.Format("2006-01-02 15:04:05")
}
func Println(str1 string, str2 string) string {
	return str1 + str2
}
func main() {
	r := gin.Default()
	r.SetFuncMap(template.FuncMap{
		"UnixToTime": UnixToTime,
		"Println":    Println,
	})
	r.LoadHTMLGlob("go_gin/server_02/templates/**/*")
	r.Static("/static/", "./static/")
	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", gin.H{
			"msg":   "我是后台数据",
			"title": "我是标题-->",
			"score": 78,
			"date":  time.Now().Unix(),
		})
	})
	r.GET("/new", func(c *gin.Context) {
		c.HTML(200, "news.html", gin.H{
			"msg":   "我是后台数据",
			"title": "我是标题-->",
			"score": 78,
			"date":  time.Now().Unix(),
		})
	})
	r.GET("/date", func(c *gin.Context) {
		username := c.Query("username")
		age := c.Query("age")
		page := c.DefaultQuery("page", "1")
		c.JSON(200, gin.H{
			"username": username,

			"age":  age,
			"page": page,
		})
	})
	r.GET("/login", func(c *gin.Context) {
		c.HTML(200, "login.html", "")

	})
	r.GET("/userLogin", func(c *gin.Context) {
		/*
			username := c.PostForm("username")
			password := c.PostForm("password")
			age := c.PostForm("age")
			page := c.DefaultPostForm("page", "1")
		*/
		user := &UserInfo{}
		//绑定到结构体
		if err := c.ShouldBind(user); err == nil {
			c.JSON(200, user)
		} else {
			c.JSON(200, gin.H{
				"err": err.Error(),
			})
		}
	})
	//获取xml数据
	r.POST("/xml", func(c *gin.Context) {
		data, _ := c.GetRawData() //获取c.REquest.Body 读取请求数据
		art1 := &Article{}
		if err := xml.Unmarshal(data, &art1); err == nil {
			c.JSON(200, art1)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"err": err.Error(),
			})
		}
	})

	//动态路由
	// 可以直接在路径后面跟参数，带上冒号就能识别，不用非得加斜杠
	r.GET("/param/:cid:pwd:haha", func(c *gin.Context) {
		cid := c.Param("cid")
		c.String(200, cid)
	})
	r.Run(":8888")
}
