package model

import (
	"chatroom/common/message"
	"encoding/json"
	"fmt"
	"github.com/garyburd/redigo/redis"
)

//第一一个全局变量，在服务期开始运行时就初始化连接池
var (
	MyUserDao UserDao
)

//负责与redis操作，结构体中要有链接池
type UserDao struct {
	pool *redis.Pool
}

//工厂模式创建UserDao实例
func NewUserDao(pool *redis.Pool) (userDao *UserDao) {

	userDao = &UserDao{
		pool: pool,
	}
	return
}

//提供的方法

//提供什么方法
// 1.根据用户id返回user的实例+err
func (ud *UserDao) getUserById(c redis.Conn, id int) (user *User, err error) {
	//通过给定的连接池去redis查询该用户
	res, err := redis.String(c.Do("hget", "users", id))
	if err != nil {
		//错误
		if err == redis.ErrNil {
			err = Error_USER_NOTEXISTS
		}
		return
	}
	user = &User{}
	//将res反序列化一个User实例
	err = json.Unmarshal([]byte(res), user)
	if err != nil {
		fmt.Println("fan序列化失败", err)
		return
	}
	return
}

//完成登录校验Login
//1Login完成对用户的验证
//2如果用户的Id和Pwd都正确的话则返回User实例
//3如果用户的ID或密码错误，则返回对应的错误信息
func (ud *UserDao) Login(userId int, userPwd string) (user *User, err error) {
	//先从UserDao里取出一个链接
	c := ud.pool.Get()
	defer c.Close()
	user, err = ud.getUserById(c, userId)
	if err != nil {
		return
	}
	//这是证明用户获取到了
	if user.UserPwd != userPwd {
		err = Error_USER_PWD
		return
	}
	return
}
func (ud *UserDao) Register(user *message.User) (err error) {
	//先从UserDao里取出一个链接
	c := ud.pool.Get()
	defer c.Close()
	_, err = ud.getUserById(c, user.UserId)
	if err == nil {
		err = Error_USER_EXISTS
		return
	}
	//这是说明这个Id还没有进行注册，则可以完成注册
	data, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
		return
	}

	//入库
	_, err = c.Do("hset", "users", user.UserId, string(data))
	if err != nil {
		fmt.Println("注册入库失败:", err)
		return
	}
	return
}
