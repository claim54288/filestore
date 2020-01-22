package redis

import (
	"time"

	"github.com/garyburd/redigo/redis"
)

var (
	pool      *redis.Pool
	redisHost = "192.168.159.131:6379"
	redisPass = "" //密码，我没设
)

func newRedisPool() *redis.Pool {
	return &redis.Pool{
		Dial: func() (redis.Conn, error) {
			//打开链接
			c, err := redis.Dial("tcp", redisHost)
			if err != nil {
				panic(err)
			}
			//访问认证
			if _, err := c.Do("AUTH", redisPass); err != nil {
				c.Close()
				panic(err)
			}
			return c, nil
		},
		TestOnBorrow: func(conn redis.Conn, t time.Time) error { //一个用来检测链接健康状态的方法，返回报错这个链接就会被关闭
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := conn.Do("PING")
			return err
		},
		MaxIdle:     100,
		MaxActive:   50,
		IdleTimeout: 5 * time.Minute,
	}

}

func init() {
	pool = newRedisPool()
}

func RedisPool() *redis.Pool {
	return pool
}
