package controllers

import (
	"bookstore/dao"
	"bookstore/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
}

func (UserHandler) Index(c *gin.Context) {
	pageNo := c.Query("pageNo")
	minPrice := c.Query("min")
	maxPrice := c.Query("max")
	if pageNo == "" {
		pageNo = "1"
	}
	page := &models.Page{}
	if minPrice == "" && maxPrice == "" {
		//没传最大最小的价格就调用获取书籍
		page, _ = dao.GetPageBooks(pageNo)
	} else {
		// 传了价格就调用价格
		page, _ = dao.GetPageBooksByPrice(pageNo, minPrice, maxPrice)
		//将价格设置到结构体实例中
		page.MinPrice = minPrice
		page.MaxPrice = maxPrice
	}
	flag, session := dao.IsLogin(c)
	if flag {
		//已经登录
		//设置page中login字段和username字段
		page.IsLogin = true
		page.UserName = session.Username
	}
	c.HTML(http.StatusOK, "index.html", page)
}
func (UserHandler) Login(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", "")
}
func (UserHandler) Regist(c *gin.Context) {
	c.HTML(http.StatusOK, "regist.html", "")
}
