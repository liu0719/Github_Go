package dao

import (
	"go_web/bookstore/model"
	"go_web/bookstore/utils"
	"net/http"
)

// AddSession 向数据库中添加Session
func AddSession(sess *model.Session) error {
	sql := "insert into sessions values(?,?,?)"
	_, err := utils.Db.Exec(sql, sess.SessionId, sess.UserName, sess.UserId)
	if err != nil {
		return err
	}
	return nil
}

// DeleteSession 向数据库中添加Session
func DeleteSession(sessId string) error {
	sql := "delete from sessions where session_id = ?"
	_, err := utils.Db.Exec(sql, sessId)
	if err != nil {
		return err
	}
	return nil
}

// 根据SessionId值获取Session
func GetSession(sessId string) (*model.Session, error) {
	sql := "select session_id,username,user_id from sessions where session_id = ?"
	// 预编译
	inStmt, err := utils.Db.Prepare(sql)
	if err != nil {
		return nil, err
	}

	//执行
	row := inStmt.QueryRow(sessId)
	sess := &model.Session{}
	// 扫描数据库中字段值是Session的字段赋值
	row.Scan(&sess.SessionId, &sess.UserName, &sess.UserId)
	return sess, err
}

// 判断用户是否登录
func IsLogin(r *http.Request) (bool, *model.Session) {
	cookie, _ := r.Cookie("user")
	if cookie != nil {
		//获取到cookie
		//不为空就获取cookie的值
		cookieValue := cookie.Value

		//去数据库中查询对应的Session
		session, _ := GetSession(cookieValue)
		if session.UserId > 0 {
			//已经登录
			return true, session
		}
	}
	//没有登陆
	return false, nil

}
