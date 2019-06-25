package main

import (
	"github.com/redigo/redis"
	"time"
)

//定义一个全局的pool
var pool redis.Pool

func initPool(address string, maxIdle, maxActive int, idleTimeout time.Duration) {

	pool = redis.Pool{
		MaxIdle:     maxIdle,     //最大空闲链接数
		MaxActive:   maxActive,   //表示和数据库的最大连接数,0表示没有限制
		IdleTimeout: idleTimeout, //最大空闲时间
		Dial: func() (conn redis.Conn, e error) { //初始化连接代码,连接哪个ip的redis
			return redis.Dial("tcp", address)
		},
	}
}
