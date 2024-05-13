package utils

import (
	"database/sql"
	"errors"
	_ "github.com/go-sql-driver/mysql"
)

var (
	Db  *sql.DB
	err error
)

func init() {
	Db, err = sql.Open("mysql", "root:123456@tcp(localhost:3306)/test")
	if err != nil {
		panic(errors.New("链接错误"))
	}
}
