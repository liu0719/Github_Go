package dao

import (
	"fmt"
	"go_web/bookstore/model"
	"go_web/bookstore/utils"
)

//根据用户名和密码从数据库中查询记录
func CheckUserNameAndPassword(username string,
	password string) (*model.User, error) {
	// sql语句
	sqlStr := "select id,username,password,email from users where username = ? and password = ?"

	//执行
	// QueryRow只查一条
	row := utils.Db.QueryRow(sqlStr, username, password)
	user := &model.User{}
	row.Scan(&user.Id, &user.Username, &user.Password, &user.Email)
	return user, nil
}

//根据用户名和密码从数据库中查询记录
func CheckUserName(username string) (*model.User, error) {
	// sql语句
	sqlStr := "select id,username,password,email from users where username = ?"

	//执行
	// QueryRow只查一条
	row := utils.Db.QueryRow(sqlStr, username)
	user := &model.User{}
	row.Scan(&user.Id, &user.Username, &user.Password, &user.Email)
	return user, nil
}

//SaveUser 向数据库中插入信息
func SaveUser(username string, password string, email string) error {
	sql := "insert into users(username,password,email) values(?,?,?)"
	//执行
	_, err := utils.Db.Exec(sql, username, password, email)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
