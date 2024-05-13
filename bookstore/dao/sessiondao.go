package dao

import (
	"bookstore/models"
	"bookstore/utils"

	"github.com/gin-gonic/gin"
)

// 判断是否登录
func IsLogin(c *gin.Context) (bool, *models.Session) {
	cookie, _ := c.Cookie("user")
	if cookie != "" {
		//获取到cookie
		//不为空就获取cookie的值
		cookieValue := cookie

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

// 根据Session的idsession值获取
func GetSession(sessionId string) (*models.Session, error) {
	sessionQuery := &models.SessionQuery{}
	result := utils.DB.Select("session_id,username,user_id").Where("session_id = ?", sessionId).First(sessionQuery)
	if result.Error != nil {
		return nil, result.Error
	}
	// 临时的query和真Session进行交换
	session := &models.Session{
		SessionId: sessionQuery.SessionId,
		Username:  sessionQuery.Username,
		UserId:    sessionQuery.UserId,
	}
	return session, nil
}

// 向数据库中添加session
func AddSession(session *models.Session) error {
	sessionQuery := &models.SessionQuery{
		SessionId: session.SessionId,
		Username:  session.Username,
		UserId:    session.UserId,
	}
	result := utils.DB.Create(sessionQuery)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// 根据SessionId删除Session
func DeleteSession(sessionId string) error {
	sessionQuery := &models.SessionQuery{
		SessionId: sessionId,
	}
	result := utils.DB.Where("session_id=?", sessionId).Delete(sessionQuery)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
