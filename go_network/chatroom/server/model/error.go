package model

import "errors"

var (
	Error_USERR_NOTEXISTS = errors.New("该账号不存在，请注册后使用")
	Error_USERR_EXISTS    = errors.New("该用户已经存在,请更换账号注册或登录")
	Error_USER_PWD        = errors.New("密码不正确,请重试")
)
