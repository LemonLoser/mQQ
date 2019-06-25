package dao

import (
	"encoding/json"
	"fmt"
	"github.com/redigo/redis"
	"gocode/mQQ/common"
	"gocode/mQQ/server/model"
)

//在服务器启动后实例化一个userDao的实例
//做成一个全局变量,需要和redis操作时直接使用即可
var MyUserDao *UserDao

//定义一个UserDao结构体,并对User对象进行curd的操作,这里curd操作要进行封装
type UserDao struct {
	pool redis.Pool
}

//使用工厂模式创建一个UserDao实例
func NewUserDao(pool redis.Pool) (userdao *UserDao) {
	userdao = &UserDao{
		pool: pool,
	}
	return
}

//用户注册
func (this *UserDao) Register(user common.User) (err error) {
	conn := this.pool.Get()
	defer conn.Close()
	//先从数据库中查找是否已存在该用户
	_, err = this.Search(conn, user.UserId)
	if err == nil { //err不为空说明该用户已存在
		err = model.ERROR_USER_EXIST
		return
	}
	//序列化user结构体
	data, err := json.Marshal(user)
	if err != nil {
		return
	}
	fmt.Println("即将存入数据库的data值", string(data))
	//存入数据库
	_, err = conn.Do("HSet", "user", fmt.Sprintf("%d", user.UserId), string(data))
	if err != nil {
		fmt.Println("入库出错!!!")
		return
	}
	return
}

//通过用户Id查询用户
func (this *UserDao) Search(conn redis.Conn, userid int) (user model.User, err error) {
	fmt.Println("传入的userid", userid)
	res, err := redis.String(conn.Do("HGet", "user", fmt.Sprintf("%d", userid)))
	fmt.Printf("res的类型%T,res的值%v\n", res, res)
	if err != nil {
		if err == redis.ErrNil {
			err = model.ERRER_USER_NOTEXIST
			return
		}
		return
	}
	//反序列化res成结构体
	//声明一个user
	user = model.User{}
	err = json.Unmarshal([]byte(res), &user)
	if err != nil {
		fmt.Println("user struct unmarshal fail err=", err)
		return
	}
	fmt.Println("从数据库得到的user=", user)
	return
}

//登陆校验
//如果用户名密码都正确则返回一个用户实例
//如果有错则返回对应错误信息
func (this *UserDao) Login(userId int, userPwd string) (user model.User, err error) {
	fmt.Println("已经执行Search 和 Check方法")
	//从连接池中取出一个连接
	conn := this.pool.Get()
	defer conn.Close()
	user, err = this.Search(conn, userId)
	fmt.Println("for checked user=", user)
	if err != nil {
		return
	}
	//用户已经获取到了
	//校验密码
	if user.UserPwd != userPwd {
		err = model.ERROR_USER_PWD
		return
	}
	return
}
