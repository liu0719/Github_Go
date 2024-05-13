package main

import (
	"net/http"
	"text/template"
)

type User struct {
	Id       int
	Name     string
	Password string
	Email    string
}

// 条件动作
func testIf(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("web03_action/index.html")
	age := 17
	t.Execute(w, age < 18)
}

// 迭代动作
func testRange(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("web03_action/range.html")
	var users []*User
	user1 := &User{
		Id:       1,
		Name:     "小明",
		Password: "123",
		Email:    "123@qq.com",
	}
	users = append(users, user1)
	user2 := &User{
		Id:       2,
		Name:     "小红",
		Password: "1ffwpf",
		Email:    "123@154.com",
	}
	users = append(users, user2)
	user3 := &User{
		Id:       3,
		Name:     "张三",
		Password: "3ti3",
		Email:    "123@wfwfq.com",
	}
	users = append(users, user3)
	t.Execute(w, users)
}

//包含动作
func testTemplate(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("web03_action/template1.html", "web03_action/template2.html")
	t.Execute(w, "我能在两个文件中显示吗")
}

// 设置动作
func testWith(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("web03_action/with.html")
	t.Execute(w, "狸猫")
}

//定义动作
func testDefine(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("web03_action/define.html")
	t.ExecuteTemplate(w, "model", "")
}

//定义动作2
func testDefine2(w http.ResponseWriter, r *http.Request) {
	age := 20
	var t *template.Template
	if age > 18 {
		t, _ = template.ParseFiles("web03_action/define2.html", "web03_action/content1.html")
	} else {
		t, _ = template.ParseFiles("web03_action/define2.html", "web03_action/content2.html")
	}

	t.ExecuteTemplate(w, "model", "")
}
func main() {
	http.HandleFunc("/if", testIf)
	http.HandleFunc("/define", testDefine)
	http.HandleFunc("/define2", testDefine2)
	http.HandleFunc("/range", testRange)
	http.HandleFunc("/with", testWith)
	http.HandleFunc("/template", testTemplate)
	http.ListenAndServe(":9999", nil)
}
