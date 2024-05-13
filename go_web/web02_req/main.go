package main

import (
	"encoding/json"
	"fmt"
	"go_web/web01_db/model"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// 获取请求行的信息
	//获取访问路径
	fmt.Fprintln(w, "请求地址是:", r.URL.Path)
	//获取请求的字符串集
	fmt.Fprintln(w, "请求地址后的字符串是:", r.URL.RawQuery)
	fmt.Fprintln(w, "请求头的所有信息:", r.Header)
	fmt.Fprintln(w, "请求头的Accept信息:", r.Header["Accept"])
	//get方法获取到的是没有中括号的值
	fmt.Fprintln(w, "请求头的Accept信息:", r.Header.Get("Accept"))
	//获取请求体的长度
	len := r.ContentLength
	//创建一个切片
	body := make([]byte, len)
	//将请求体的内容读到body哩,会把请求体的内容读走，不会再保留
	// r.Body.Read(body)
	fmt.Fprintln(w, "请求体的内容:", string(body))

	//解析表单
	r.ParseForm()
	//获取请求参数
	fmt.Fprintln(w, "请求参数有:", r.Form)
	fmt.Fprintln(w, "只获取Post请求的参数", r.PostForm)
	fmt.Fprintln(w, "获取单个username", r.FormValue("username"))
}

func testJsonRes(w http.ResponseWriter, r *http.Request) {
	//设置响应的类型
	w.Header().Set("Content-Type", "application/json")
	//创建一个User
	user := model.Users{
		Id:       100,
		Username: "小明",
		Password: "122443",
		Email:    "wwjd@qq.com",
	}
	data, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
		return
	}
	//响应给客户端
	w.Write(data)
	// 重定向
}
func testRedire(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Location", "http://www.baidu.com")
	w.WriteHeader(302)
}
func main() {
	http.HandleFunc("/hello", handler)
	http.HandleFunc("/jsonres", testJsonRes)
	http.HandleFunc("/redire", testRedire)

	http.ListenAndServe(":8888", nil)
}
