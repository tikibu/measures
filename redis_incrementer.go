package measures

import (
	"github.com/gomodule/redigo/redis"
	"strconv"
	"time"
)

type RedisIncrementer struct {
	pool *redis.Pool
}

func NewRedisIncrementer(host string, port int) (fi *RedisIncrementer) {
	dsn := host + ":" + strconv.Itoa(port)
	fi = &RedisIncrementer{
		pool: dialRedis(dsn),
	}
	return
}

func dialRedis(host string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     5,
		IdleTimeout: 0,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial(
				"tcp",
				host,
				redis.DialConnectTimeout(time.Second*10),
			)
			if err != nil {
				return nil, err
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}

func (fi *RedisIncrementer) IncAndGet(key string, by float64) (res float64, err error) {
	conn := fi.pool.Get()
	defer conn.Close()
	conn.Send("INCRBYFLOAT", key, by)
	conn.Flush()
	res, err = redis.Float64(conn.Receive())
	return
}

func (fi *RedisIncrementer) Get(key string) (res float64, err error) {
	conn := fi.pool.Get()
	defer conn.Close()
	conn.Send("GET", key)
	conn.Flush()
	raw, err := conn.Receive()
	if err == nil && raw == nil {
		res = 0.0
		return
	}
	res, err = redis.Float64(raw, err)
	return
}
