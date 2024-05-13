package handler

import (
	"fmt"
	"net/http"
	"news_0417/dao"
	"news_0417/model"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
}

func (UserHandler) Index(c *gin.Context) {
	news, err := dao.GetNews()
	if err != nil {
		fmt.Println(err)
		return
	}
	c.HTML(http.StatusOK, "index.html", news)
}
func (UserHandler) AddNewsPage(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		c.HTML(http.StatusOK, "addnew.html", "")
	}
}
func (UserHandler) AddNews(c *gin.Context) {
	id := c.PostForm("id")
	title := c.PostForm("title")
	author := c.PostForm("author")
	content := c.PostForm("content")
	class := c.PostForm("class")
	iClass, _ := strconv.Atoi(class)
	iId, _ := strconv.Atoi(id)
	new := &model.News{
		Id:         iId,
		Title:      title,
		Author:     author,
		Context:    content,
		ClassId:    iClass,
		CreateTime: time.Now().Format("2006-01-02 15:04:05"),
	}
	if id == "" {
		err := dao.AddNews(new)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		err := dao.UpdateNews(new)
		if err != nil {
			fmt.Println(err)
		}
	}
	c.Redirect(302, "/")
}
func (UserHandler) DeleteNew(c *gin.Context) {
	id := c.Query("id")
	err := dao.DeleteNew(id)
	if err != nil {
		fmt.Println(err)
	}
	c.Redirect(302, "/")
}
func (UserHandler) UpdateNew(c *gin.Context) {
	id := c.Query("id")
	new, err := dao.GetNewsById(id)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(new)
	c.HTML(http.StatusOK, "addnew.html", new)
}
func (UserHandler) SelectNewsByAuhtor(c *gin.Context) {
	author := c.PostForm("author")
	news, err := dao.GetNewsByAuthor(author)
	if err != nil {
		fmt.Println(err)
	}
	c.HTML(http.StatusOK, "index.html", news)

}
func (UserHandler) SelectNewsByClass(c *gin.Context) {
	classId := c.PostForm("class")
	if classId == "0" {
		news, err := dao.GetNews()
		if err != nil {
			fmt.Println(err)
		}
		c.HTML(http.StatusOK, "index.html", news)
		return
	}
	news, err := dao.GetNewsByClass(classId)
	if err != nil {
		fmt.Println(err)
	}
	c.HTML(http.StatusOK, "index.html", news)

}
