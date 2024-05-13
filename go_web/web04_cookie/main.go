package main

import (
	"net/http"
)

//SetCookie 添加cookie
func setCookie(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{
		Name:     "user",
		Value:    "admin",
		HttpOnly: false,
		MaxAge:   -1,
	}
	cookie2 := http.Cookie{
		Name:     "user2",
		Value:    "admin2",
		HttpOnly: false,
		MaxAge:   -10000,
	}
	// 将cookie发送给浏览器
	w.Header().Set("Set-Cookie", cookie.String())
	// 添加第二个cookie
	w.Header().Add("Set-cookie", cookie2.String())
	// http.SetCookie(w, &cookie)
	// cookiestr := r.Header.Get("cookie")
	// fmt.Println(cookiestr)

}

func GetCookies(w http.ResponseWriter, r *http.Request) {
	// 获取所有的cookie
	// cookies := r.Header["Cookie"]

	//直接获取单个,
	cookie2 := &http.Cookie{}
	cookie2.MaxAge = -10000
	http.SetCookie(w, cookie2)
	// fmt.Println("得到的所有cookie", cookie2)

}

func main() {
	http.HandleFunc("/", setCookie)
	http.HandleFunc("/getcookies", GetCookies)

	http.ListenAndServe(":9999", nil)
}
