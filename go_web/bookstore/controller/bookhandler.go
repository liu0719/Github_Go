package controller

import (
	"go_web/bookstore/dao"
	"go_web/bookstore/model"
	"net/http"
	"strconv"
	"text/template"
)

// 首页处理器
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	pageNo := r.FormValue("pageNo")
	if pageNo == "" {
		pageNo = "1"
	}
	page, err := dao.GetPageBooks(pageNo)
	if err != nil {
		return
	}
	//解析模板
	t := template.Must(template.ParseFiles("bookstore/views/index.html"))
	//执行
	t.Execute(w, page)
}

// 获取图书
// func GetBooks(w http.ResponseWriter, r *http.Request) {
// 	//调用BookDao中获取所有图书的函数
// 	books, _ := dao.GetBooks()
// 	// 解析模板
// 	t := template.Must(template.ParseFiles("bookstore/views/pages/manager/book_manager.html"))
// 	// 执行
// 	t.Execute(w, books)
// }

//添加图书
// func AddBook(w http.ResponseWriter, r *http.Request) {
// 	title := r.PostFormValue("title")
// 	author := r.PostFormValue("author")
// 	price := r.PostFormValue("price")
// 	sales := r.PostFormValue("sales")
// 	stock := r.PostFormValue("stock")
// 	// string转换float64等
// 	fprice, _ := strconv.ParseFloat(price, 64)
// 	isales, _ := strconv.Atoi(sales)
// 	istock, _ := strconv.Atoi(stock)
// 	// 创建book

// 	book := &model.Book{
// 		Title:   title,
// 		Price:   fprice,
// 		Author:  author,
// 		Stock:   istock,
// 		Sales:   isales,
// 		ImgPath: "/static/img/default.jpg",
// 	}
// 	err := dao.AddBook(book)
// 	if err != nil {
// 		panic(err)
// 	}

// 	//添加后再查询一次数据库，更新数据
// 	GetBooks(w, r)
// }

// 删除图书
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	// 获取要删除的图书Id
	bookIdStr := r.FormValue("Id")
	bookId, err := strconv.Atoi(bookIdStr)
	if err != nil {
		return
	}
	//调用bookDao中要删除的函数
	dao.DeleteBook(bookId)
	//删除后再查询一次数据库，更新数据
	GetPageBooks(w, r)
}

// 去更新或添加图书页面
func ToUpdateBookPage(w http.ResponseWriter, r *http.Request) {
	// 获取要更新的图书Id
	bookId := r.FormValue("bookId")
	book, err := dao.GetBookById(bookId)
	if book.Id > 0 {
		//更新图书y页面
		// 解析模板
		t := template.Must(template.ParseFiles("bookstore/views/pages/manager/book_edit.html"))
		// 执行
		t.Execute(w, book)
	} else {
		//添加图书页面
		// 解析模板
		t := template.Must(template.ParseFiles("bookstore/views/pages/manager/book_edit.html"))
		// 执行
		t.Execute(w, "")
	}
	if err != nil {
		return
	}

}

// 更新或添加图书
func UpdateOrAddBook(w http.ResponseWriter, r *http.Request) {
	id := r.PostFormValue("bookId")
	title := r.PostFormValue("title")
	author := r.PostFormValue("author")
	price := r.PostFormValue("price")
	sales := r.PostFormValue("sales")
	stock := r.PostFormValue("stock")
	// string转换float64等
	fprice, _ := strconv.ParseFloat(price, 64)
	isales, _ := strconv.Atoi(sales)
	istock, _ := strconv.Atoi(stock)
	iId, _ := strconv.Atoi(id)
	// 创建book

	book := &model.Book{
		Id:      iId,
		Title:   title,
		Price:   fprice,
		Author:  author,
		Stock:   istock,
		Sales:   isales,
		ImgPath: "/static/img/default.jpg",
	}
	if book.Id > 0 {
		//更新图书
		err := dao.UpdateBook(book)
		if err != nil {
			panic(err)
		}
	} else {
		// 添加图书
		dao.AddBook(book)
	}

	//添加后再查询一次数据库，更新数据
	GetPageBooks(w, r)
}

// 获取带分页图书
func GetPageBooks(w http.ResponseWriter, r *http.Request) {
	// 获取页码
	pageNo := r.FormValue("pageNo")
	if pageNo == "" {
		pageNo = "1"
	}
	page, err := dao.GetPageBooks(pageNo)
	if err != nil {
		return
	}
	// 解析模板
	t := template.Must(template.ParseFiles("bookstore/views/pages/manager/book_manager.html"))
	// 执行
	t.Execute(w, page)
}

// 根据价格获取图书
func GetPageBooksByPrice(w http.ResponseWriter, r *http.Request) {
	// 获取页码
	pageNo := r.FormValue("pageNo")
	minPrice := r.FormValue("min")
	maxPrice := r.FormValue("max")

	if pageNo == "" {
		pageNo = "1"
	}
	var page *model.Page
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

	// 调用函数IsLogin判断是否已经登录
	flag, session := dao.IsLogin(r)
	if flag {
		//已经登录
		//设置page中的login和UserName字段
		page.IsLogin = true
		page.UserName = session.UserName
	}

	// 解析模板
	t := template.Must(template.ParseFiles("bookstore/views/index.html"))
	// 执行
	t.Execute(w, page)
}
