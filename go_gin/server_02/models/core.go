package models

import (
	"fmt"

	"gorm.io/driver/mysql"
	_ "gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func init() {
	//https://gorm.io/zh_CN/docs/connecting_to_the_database.html
	dsn := "root:123456@tcp(localhost:3306)/gin_demo?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
}
