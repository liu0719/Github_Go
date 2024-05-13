package controller

import (
	"go_web/bookstore/dao"
	"go_web/bookstore/model"
	"go_web/bookstore/utils"
	"net/http"
	"os"
	"text/template"
)

// Logout 登出
func Logout(w http.ResponseWriter, r *http.Request) {
	//获取cookie的Value值
	cookie, err := r.Cookie("user")
	if err != nil {
		return
	}
	//获取cookie的Value
	cookieValue := cookie.Value
	//删除数据库中与之对应的Session
	err = dao.DeleteSession(cookieValue)
	if err != nil {
		return
	}
	//让cookie失效,小于零相当于删除cookie
	cookie.MaxAge = -1
	http.SetCookie(w, cookie)
	//登出后去首页
	GetPageBooksByPrice(w, r)
}

// Login 处理登录请求
func Login(w http.ResponseWriter, r *http.Request) {
	//判断当前用户是否已经登录
	flag, _ := dao.IsLogin(r)
	if flag {
		//已经登录
		GetPageBooksByPrice(w, r)
		return
	}
	//获取用户名和密码
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	//调用userdao验证用户名和密码
	user, _ := dao.CheckUserNameAndPassword(username, password)
	if user.Id > 0 {
		// 正确
		//使用UUId生成SessionID
		uuid := utils.CreateUUID()
		// 创建一个Session
		sess := &model.Session{
			SessionId: uuid,
			UserName:  user.Username,
			UserId:    user.Id,
		}
		//将该Session保存到数据库中
		dao.AddSession(sess)
		// 创建cookie，与Session相关联
		cookie := http.Cookie{
			Name:     "user",
			Value:    uuid,
			HttpOnly: true,
		}
		//将该cookie发送给浏览器
		http.SetCookie(w, &cookie)

		t := template.Must(template.ParseFiles("bookstore/views/pages/user/login_success.html"))
		t.Execute(w, user)
	} else {
		//用户名或密码不正确
		t := template.Must(template.ParseFiles("bookstore/views/pages/user/login.html"))
		t.Execute(w, "用户名或密码不正确!")
	}
}

// Register 处理登录请求
func Register(w http.ResponseWriter, r *http.Request) {
	//获取用户名和密码
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	email := r.PostFormValue("email")
	//调用userdao验证用户名和密码
	user, _ := dao.CheckUserName(username)
	if user.Id > 0 {
		//用户名已经存在
		t := template.Must(template.ParseFiles("bookstore/views/pages/user/regist.html"))
		t.Execute(w, "用户名已存在！")
	} else {
		//用户名可用，保存到数据库中
		dao.SaveUser(username, password, email)
		// 跳转到注册成功页面
		t := template.Must(template.ParseFiles("bookstore/views/pages/user/regist_success.html"))
		t.Execute(w, "")
	}
}

func CheckUserName(w http.ResponseWriter, r *http.Request) {
	//获取用户输入的用户名
	username := r.PostFormValue("username")
	// 调用userdao验证用户名和密码
	user, _ := dao.CheckUserName(username)
	if user.Id > 0 {
		//用户名已经存在
		w.Write([]byte("用户名已存在！"))
	} else {
		w.Write([]byte("<font style=\"color:green;\">用户名可用</font>"))
	}
}
func Favicon(w http.ResponseWriter, r *http.Request) {
	date, err := os.ReadFile("bookstore/views/favicon.ico")
	if err != nil {
		return
	}
	w.Write(date)
}
