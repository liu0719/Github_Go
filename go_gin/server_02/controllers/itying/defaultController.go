package itying

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type DefaultController struct {
}

func (con DefaultController) Index(c *gin.Context) {
	//设置session
	// 获取session对象
	session := sessions.Default(c)
	// 设置Session
	session.Set("username", "张三111")
	// 保存Session,一定要保存之后才能在别的页面访问
	session.Save() // 必须调用
	/*
		//设置cookie
		// 3600是秒
		//超过设置的时长就会过期
		c.SetCookie("username", "张三", 3600, "/", ".haha.com", false, true)
	*/
	c.HTML(200, "index.html", gin.H{
		"msg":   "杨佳鑫大傻b",
		"date":  time.Now().Unix(),
		"title": "标题一号",
	})
}
func (con DefaultController) News(c *gin.Context) {
	/*
		// 获取cookie
		username, _ := c.Cookie("username")
		c.String(200, "新闻----cookie"+username)
	*/
	// 获取session
	session := sessions.Default(c)
	//设置session的过期时间
	session.Options(sessions.Options{
		MaxAge: 3600 * 6, //6小时（单位是秒）
	})
	username := session.Get("username")
	c.String(http.StatusOK, username.(string))
}
func (con DefaultController) Shop(c *gin.Context) {
	// 获取cookie
	username, _ := c.Cookie("username")
	c.String(200, "新闻----cookie"+username)
	fmt.Println(username)
}
func (con DefaultController) Delete(c *gin.Context) {
	//删除cookie就是需要从新设置
	//maxAge设置成负数就表示删除cookie
	c.SetCookie("username", "张三", -1, "/", "localhost", false, true)
	c.SetCookie("username", "张三", -1, "/", "10.50.142.216", false, true)
	c.String(200, "删除成功")
}
