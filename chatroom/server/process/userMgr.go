package process

import (
	"fmt"
)

var (
	userMgr *UserMgr
)

type UserMgr struct {
	onlineUsers map[int]*UserProcess
}

//完成初始化工作
func init() {
	userMgr = &UserMgr{
		onlineUsers: make(map[int]*UserProcess, 1024),
	}
}

//完成对onlineUser的添加
func (this *UserMgr) AddOnlineUser(up *UserProcess) {
	this.onlineUsers[up.UserId] = up
}

//删除功能
func (this *UserMgr) DelOnlineUser(userId int) {
	delete(this.onlineUsers, userId)
}

//查询
func (this *UserMgr) GetAllOnlineUser() map[int]*UserProcess {
	return this.onlineUsers
}

//根据Id返回对应的值
func (this *UserMgr) GetUserById(userId int) (up *UserProcess, err error) {
	//从map中取出一个值
	up, ok := this.onlineUsers[userId]
	if !ok { //说明查找得用晃前不在线
		err = fmt.Errorf("用户%v不在线。。。", userId)
		return
	}
	return
}
