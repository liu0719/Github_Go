package middlewares

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// 中间件
func InitIndex(c *gin.Context) {
	fmt.Println("我是中间件---1")
	//中间件和控制器数据共享，创建一个数据供其他中间件和控制器获取和使用，
	// 必须执行完该条语句之后其他中间件和控制器藏看见
	c.Set("username", "刘广硕")
	//该方法会跳过函数剩余部分，直接执行下一个处理函数，等下一个函数处理完返回之后才会就须执行剩余部分
	// 协程
	cCp := c.Copy()
	go func() {
		//必须使用只读副本

		time.Sleep(time.Second * 5)

		//这里创建副本
		fmt.Println("Done! in path " + cCp.Request.URL.Path)
	}()
	c.Next()
	// 终止剩余请求函数，但会执行完当前函数
	// c.Abort()
	fmt.Println("我是中间件---2")

	fmt.Println(c.Request.URL)
	fmt.Println(time.Now())

}
