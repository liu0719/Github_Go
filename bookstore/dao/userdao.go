package dao

import (
	"bookstore/models"
	"bookstore/utils"
)

// 根据用户名和密码查询记录,判断是否登录
func CheckUserNameAndPassword(username, password string) (*models.User, error) {
	user := &models.User{}

	result := utils.DB.Select("id,username,password,email").Where("username=? and password=?", username, password).First(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}
func CheckUserName(username string) (*models.User, error) {
	user := &models.User{}
	result := utils.DB.Select("id,username,password,email").Where("username = ? ", username).First(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}
func SaveUser(username, password, email string) error {
	user := &models.User{
		Username: username,
		Password: password,
		Email:    email,
	}
	result := utils.DB.Create(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
