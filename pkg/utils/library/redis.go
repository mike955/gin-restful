package library

import (
	"gin-restful/pkg/setting"
	"github.com/gomodule/redigo/redis"
	"time"
)

type RedisConn struct {
	Rdb *redis.Pool
}

var RedisClient *redis.Pool

func RedisSetup() error {
	RedisClient = &redis.Pool{
		MaxIdle:     setting.RedisSetting.MaxIdle,
		MaxActive:   setting.RedisSetting.MaxActive,
		IdleTimeout: setting.RedisSetting.IdleTimeout,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", setting.RedisSetting.Host)
			if err != nil {
				return nil, err
			}
			if setting.RedisSetting.Password != "" {
				if _, err := c.Do("AUTH", setting.RedisSetting.Password); err != nil {
					c.Close()
					return nil, err
				}
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}

	return nil
}

func NewRedis() *RedisConn {
	rdb := &RedisConn{Rdb:RedisClient}
	return rdb
}

func (rdb *RedisConn)Exists(key string) bool {
	conn := rdb.Rdb.Get()
	defer conn.Close()

	exists, err := redis.Bool(conn.Do("EXISTS", key))
	if err != nil {
		return false
	}
	return exists
}
//
//func (rdb *RedisConn)()  {
//
//}