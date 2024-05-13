package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func testTemplate(w http.ResponseWriter, r *http.Request) {
	//解析模板
	/*
		t, err := template.ParseFiles("web02_template/index.html")
		if err != nil {
			fmt.Println(err)
			return
		}
	*/
	//通过Must函数帮我们处理异常
	t := template.Must(template.ParseFiles("web02_template/index.html", "web02_template/index2.html"))
	//执行
	//下面这个方法，只能处理首个模板，如果传入多个模板的话，只会显示第一个
	// err := t.Execute(w, "hello template"
	//第二个文件index2.html响应
	err := t.ExecuteTemplate(w, "index2.html", "我要去index2.html")
	if err != nil {
		fmt.Println(err)
		return
	}
}
func indexCss(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/css; charset=utf-8")
	t, _ := template.ParseFiles("web02_template/index.css")
	t.Execute(w, "index.css发送完成")
}
func main() {
	http.HandleFunc("/testtemplate", testTemplate)
	http.HandleFunc("/index.css", indexCss)

	http.ListenAndServe(":8888", nil)
}
