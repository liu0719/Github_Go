package model

import (
	"fmt"
	"go_web/web01_db/utils"
)

//Users 结构体
type Users struct {
	Id       int
	Username string
	Password string
	Email    string
}

func (user *Users) AddUser() error {
	sqlStr := "insert into users(username,password,email) values(?,?,?)"
	//2.预编译
	inStmt, err := utils.Db.Prepare(sqlStr)
	if err != nil {
		fmt.Println("预编译出现异常", err)
		return err
	}
	//执行
	_, err2 := inStmt.Exec("admin1", "123456", "123@qq.com")
	if err2 != nil {
		fmt.Println("执行出现异常", err2)
		return err2
	}
	return nil
}

//不用预编译
func (user *Users) AddUser2() error {
	sqlStr := "insert into users(username,password,email) values(?,?,?)"
	//执行
	_, err := utils.Db.Exec(sqlStr, "admin2", "123456", "123@qq.com")
	if err != nil {
		fmt.Println("执行出现异常", err)
		return err
	}
	return nil
}

func (user *Users) GetUserById() (*Users, error) {
	//sql语句
	sqlStr := "select id,username,password,email from users where id = ?"
	row := utils.Db.QueryRow(sqlStr, user.Id)
	//声明
	var id int
	var username string
	var password string
	var email string
	err := row.Scan(&id, &username, &password, &email)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	u := &Users{
		Id:       id,
		Username: username,
		Password: password,
		Email:    email,
	}
	return u, nil
}

//GetAllUser获取数据库中的所有记录
func (user *Users) GetAllUser() ([]*Users, error) {
	//sql语句
	sqlStr := "select id,username,password,email from users"
	//执行
	rows, err := utils.Db.Query(sqlStr)
	if err != nil {
		return nil, err
	}
	//创建返回的切片
	var users []*Users
	for rows.Next() {
		//声明
		var id int
		var username string
		var password string
		var email string
		err := rows.Scan(&id, &username, &password, &email)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		u := &Users{
			Id:       id,
			Username: username,
			Password: password,
			Email:    email,
		}
		users = append(users, u)
	}
	return users, nil
}
