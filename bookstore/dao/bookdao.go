package dao

import (
	"bookstore/models"
	"bookstore/utils"
	"strconv"
)

// 获取分页的图书信息
func GetPageBooks(pageNo string) (*models.Page, error) {
	iPageNo, err := strconv.Atoi(pageNo)
	if err != nil {
		return nil, err
	}
	//获取数据库中的图书信息
	var count int64
	result := utils.DB.Model(&[]*models.Book{}).Count(&count)
	if result.Error != nil {
		return nil, result.Error
	}
	iCount := int(count)
	//每页显示四个
	var pageSize int = 4
	//获取总页数
	var totalPageNo int
	if iCount%pageSize == 0 {
		totalPageNo = iCount / pageSize
	} else {
		totalPageNo = iCount/pageSize + 1
	}
	books := []*models.Book{}
	//获取当前页面的图书
	result = utils.DB.Select("id,title,author,price,sales,stock,img_path").Offset((iPageNo - 1) * pageSize).Limit(pageSize).Find(&books)
	if result.Error != nil {
		return nil, result.Error
	}

	// 创建Page
	page := &models.Page{
		Books:       books,
		PageNo:      iPageNo,
		PageSize:    pageSize,
		TotalPageNo: totalPageNo,
		TotalRecord: iCount,
	}
	return page, nil
}

// 根据价格获取书籍
func GetPageBooksByPrice(pageNo, minPrice, maxPrice string) (*models.Page, error) {
	iPageNo, err := strconv.Atoi(pageNo)
	if err != nil {
		return nil, err
	}

	//获取数据库中的图书信息
	var count int64
	result := utils.DB.Model(&[]*models.Book{}).Where("price between ? and ?", minPrice, maxPrice).Count(&count)
	if result.Error != nil {
		return nil, result.Error
	}
	//总记录数
	iCount := int(count)

	//设置每页显示四条记录
	var pageSize int = 4
	//获取总页数
	var totalPageNo int
	if iCount%pageSize == 0 {
		totalPageNo = iCount / pageSize
	} else {
		totalPageNo = iCount/pageSize + 1
	}
	//获取当前页码的图书
	books := []*models.Book{}
	//获取当前页面的图书
	result = utils.DB.Select("id,title,author,price,sales,stock,img_path").Offset((iPageNo - 1) * pageSize).Limit(pageSize).Find(&books)
	if result.Error != nil {
		return nil, result.Error
	}
	// 创建Page
	page := &models.Page{
		Books:       books,
		PageNo:      iPageNo,
		PageSize:    pageSize,
		TotalPageNo: totalPageNo,
		TotalRecord: iCount,
	}
	return page, nil
}

// 删除图书
func DeleteBook(bookId int) error {
	result := utils.DB.Where("id = ?", bookId).Delete(&models.Book{})
	if result != nil {
		return result.Error
	}
	return nil
}

// 向数据库中添加一本书
func AddBook(b *models.Book) error {
	result := utils.DB.Save(b)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// 更新图书
func UpdateBook(b *models.Book) error {
	result := utils.DB.Updates(b)

	if result.Error != nil {
		return result.Error
	}
	return nil
}

// 根据ID获取图书
func GetBookById(bookId string) (*models.Book, error) {
	book := &models.Book{}
	result := utils.DB.Where("id= ?", bookId).First(book)
	if result.Error != nil {
		return nil, result.Error
	}
	return book, nil
}
