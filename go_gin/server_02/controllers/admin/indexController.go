package admin

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type IndexController struct {
}

func (index IndexController) Index(c *gin.Context) {

	fmt.Println("管理员首页")
	c.String(http.StatusOK, "--用户列表--")
	//控制器和中间件的数据共享
	username, _ := c.Get("username")
	fmt.Println("我是控制器Index,获取的共享数据--->", username)
	// 类型断言
	c.String(200, username.(string))
}
