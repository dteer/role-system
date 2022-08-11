package settings

import (
	"fmt"
	"log"
	"time"

	redigo "github.com/gomodule/redigo/redis"
)

func ConnectRedis(name string) (*redigo.Pool, error) {
	redisConf := Conf.Redis[name]
	addr := fmt.Sprintf("%s:%d", redisConf.Host, redisConf.Port)
	redis := &redigo.Pool{
		MaxIdle:     10,
		IdleTimeout: 240 * time.Second,
		MaxActive:   1000, //最大连接数
		Wait:        true, //接池链接数达到上限时，会阻塞等待其他协程用完归还之后继续执行
		Dial: func() (redigo.Conn, error) {
			redisConn, err := redigo.Dial("tcp", addr)
			if err != nil {
				return nil, err
			}
			if redisConf.Password != "" {
				if _, err := redisConn.Do("AUTH", redisConf.Password); err != nil {
					redisConn.Close()
				}
			}
			if _, err := redisConn.Do("SELECT", redisConf.DB); err != nil {
				redisConn.Close()
				panic(err)
			}
			return redisConn, nil
		},
	}
	return redis, nil
}

func InitRedis() map[string]*redigo.Pool {
	var redisDB = make(map[string]*redigo.Pool)
	for redisName := range Conf.Redis {
		redis, err := ConnectRedis(redisName)
		if err != nil {
			log.Fatal(err.Error())
			continue
		}
		redisDB[redisName] = redis
	}
	return redisDB
}

// GetRedis 获取redis实例
func RedisConnect(name string) *redigo.Pool {
	redis, ok := RedisDB[name]
	if !ok {
		return nil
	}
	return redis
}
