package controllers

import (
	"bookstore/dao"
	"bookstore/models"
	"bookstore/utils"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BookHandler struct {
}

// 管理员获取图书
func (BookHandler) GetPageBooks(c *gin.Context) {
	// 获取页码
	pageNo := c.Query("pageNo")
	if pageNo == "" {
		pageNo = "1"
	}
	page, err := dao.GetPageBooks(pageNo)
	if err != nil {
		return
	}
	c.HTML(http.StatusOK, "book_manager.html", page)
}

// 根据价格获取图书
func (BookHandler) GetPageBooksByPrice(c *gin.Context) {
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
		//设置页码，防止出现跳错误
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

// 更新或添加图书
func (BookHandler) UpDateOrAddBook(c *gin.Context) {
	id := c.PostForm("bookId")
	title := c.PostForm("title")
	author := c.PostForm("author")
	price := c.PostForm("price")
	sales := c.PostForm("sales")
	stock := c.PostForm("stock")
	file, err := c.FormFile("file")
	filepath := ""
	dst := ""
	if err == nil {
		//提交文件的form表单要加上这个属性 enctype="multipart/form-data"
		filepath, dst, err = utils.SaveFile(file)
		if err == nil {
			c.SaveUploadedFile(file, dst)
		}
	}

	fprice, _ := strconv.ParseFloat(price, 64)
	isales, _ := strconv.Atoi(sales)
	istock, _ := strconv.Atoi(stock)
	iId, _ := strconv.Atoi(id)
	//创建book
	book := &models.Book{
		Id:      iId,
		Title:   title,
		Price:   fprice,
		Author:  author,
		Stock:   istock,
		Sales:   isales,
		ImgPath: filepath,
	}
	if book.Id != 0 {
		err := dao.UpdateBook(book)
		if err != nil {
			fmt.Println(err)
			return
		}
	} else {
		dao.AddBook(book)
	}
	//添加后再查询一次数据库
	BookHandler{}.GetPageBooks(c)
}

// 更新或添加图书页面
func (BookHandler) UpDateOrAddBookPage(c *gin.Context) {
	// 获取要更新的图书Id
	bookId := c.Query("bookId")
	if bookId == "" {
		//添加图书页面
		c.HTML(http.StatusOK, "book_edit.html", "")
		return
	}
	book, err := dao.GetBookById(bookId)
	if err != nil {
		fmt.Println(err)
		return
	}
	//更新图书页面
	c.HTML(http.StatusOK, "book_edit.html", book)
}

// 删除图书
func (BookHandler) DeleteBook(c *gin.Context) {
	// 获取要删除的图书Id
	bookIdStr := c.Query("Id")
	bookId, err := strconv.Atoi(bookIdStr)
	if err != nil {
		return
	}
	//调用bookDao中要删除的函数
	err = dao.DeleteBook(bookId)
	if err != nil {
		fmt.Println(err)
		return
	}
	//删除后再查询一次数据库，更新数据
	BookHandler{}.GetPageBooks(c)
}
