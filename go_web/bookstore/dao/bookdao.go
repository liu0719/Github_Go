package dao

import (
	"fmt"
	"go_web/bookstore/model"
	"go_web/bookstore/utils"
	"strconv"
)

// GetBooks
func GetBooks() ([]*model.Book, error) {
	//sql语句
	sql := "select id,title,author,price,sales,stock,img_path from books"
	//执行
	row, err := utils.Db.Query(sql)
	if err != nil {
		return nil, err
	}
	var books []*model.Book
	for row.Next() {
		book := &model.Book{}
		//从数据库中取出赋值
		row.Scan(&book.Id, &book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock, &book.ImgPath)
		books = append(books, book)
	}
	return books, nil
}

// AddBook向数据库中添加一本书
func AddBook(b *model.Book) error {
	sql := "insert into books(title,author,price,sales,stock,img_path) values(?,?,?,?,?,?)"
	// 执行
	_, err := utils.Db.Exec(sql, b.Title, b.Author, b.Price, b.Sales, b.Stock, b.ImgPath)
	if err != nil {
		return err
	}
	return nil
}

// 根据Id，删除图书
func DeleteBook(bookId int) error {
	sql := "delete from books where id = ?"
	_, err := utils.Db.Exec(sql, bookId)
	if err != nil {
		return err
	}
	return nil
}

func GetBookById(bookId string) (*model.Book, error) {
	sql := "select id,title,author,price,sales,stock,img_path from books where id = ?"
	row := utils.Db.QueryRow(sql, bookId)
	book := &model.Book{}
	//为book字段赋值
	row.Scan(&book.Id, &book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock, &book.ImgPath)
	return book, nil
}

// 根据Id更新图书信息
func UpdateBook(b *model.Book) error {
	sql := "update books set title=?,author=?,price=?,sales=?,stock=? where id=?"
	_, err := utils.Db.Exec(sql, b.Title, b.Author, b.Price, b.Sales, b.Stock, b.Id)
	if err != nil {
		fmt.Println(err)
	}
	return nil
}

// 获取分页的图书信息
func GetPageBooks(pageNo string) (*model.Page, error) {
	iPageNo, err := strconv.ParseInt(pageNo, 10, 64)
	if err != nil {
		return nil, err
	}
	//获取数据库中的图书信息
	sql := "select count(*) from books"
	var TotalRecord int64
	//执行
	row := utils.Db.QueryRow(sql)
	row.Scan(&TotalRecord)
	//设置每页显示四条记录
	var pageSize int64 = 4
	//获取总页数
	var totalPageNo int64
	if TotalRecord%pageSize == 0 {
		totalPageNo = TotalRecord / pageSize
	} else {
		totalPageNo = TotalRecord/pageSize + 1
	}
	//获取当前页码的图书
	sql2 := "select id,title,author,price,sales,stock,img_path from books limit ?,?"
	//执行
	rows, err := utils.Db.Query(sql2, (iPageNo-1)*pageSize, pageSize)
	if err != nil {
		return nil, err
	}
	var books []*model.Book
	for rows.Next() {
		book := &model.Book{}
		rows.Scan(&book.Id, &book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock, &book.ImgPath)
		books = append(books, book)
	}
	// 创建Page
	page := &model.Page{
		Books:       books,
		PageNo:      iPageNo,
		PageSize:    pageSize,
		TotalPageNo: totalPageNo,
		TotalRecord: TotalRecord,
	}
	return page, nil
}

// 根据价格获取分页的图书信息
func GetPageBooksByPrice(pageNo string, minPrice string, maxPrice string) (*model.Page, error) {
	iPageNo, err := strconv.ParseInt(pageNo, 10, 64)
	if err != nil {
		return nil, err
	}
	//获取数据库中的图书信息
	sql := "select count(*) from books where price between ? and ?"
	var TotalRecord int64
	//执行
	row := utils.Db.QueryRow(sql, minPrice, maxPrice)
	row.Scan(&TotalRecord)
	//设置每页显示四条记录
	var pageSize int64 = 4
	//获取总页数
	var totalPageNo int64
	if TotalRecord%pageSize == 0 {
		totalPageNo = TotalRecord / pageSize
	} else {
		totalPageNo = TotalRecord/pageSize + 1
	}
	//获取当前页码的图书
	sql2 := "select id,title,author,price,sales,stock,img_path from books where price between ? and ? limit ?,?"
	//执行
	rows, err := utils.Db.Query(sql2, minPrice, maxPrice, (iPageNo-1)*pageSize, pageSize)
	if err != nil {
		return nil, err
	}
	var books []*model.Book
	for rows.Next() {
		book := &model.Book{}
		rows.Scan(&book.Id, &book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock, &book.ImgPath)
		books = append(books, book)
	}
	// 创建Page
	page := &model.Page{
		Books:       books,
		PageNo:      iPageNo,
		PageSize:    pageSize,
		TotalPageNo: totalPageNo,
		TotalRecord: TotalRecord,
	}
	return page, nil
}
