package db

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/gomodule/redigo/redis"
	"gomod/config"
	"strconv"
)

var redisClient *redis.Pool

func init() {
	cfg, _ := config.Config.GetSection("redis")
	spew.Dump(cfg)
	MaxIdle, _ := strconv.Atoi(cfg["maxIdle"])
	MaxActive, _ := strconv.Atoi(cfg["maxActive"])
	Db, _ := strconv.Atoi(cfg["db"])

	redisClient = &redis.Pool{
		MaxIdle:   MaxIdle,
		MaxActive: MaxActive,
		//IdleTimeout: MaxIdleTimeout * time.Second,
		Wait: true,
		Dial: func() (redis.Conn, error) {
			con, err := redis.Dial("tcp", cfg["address"],
				redis.DialPassword(cfg["Auth"]),
				redis.DialDatabase(Db),
				//redis.DialConnectTimeout(timeout*time.Second),
				//redis.DialReadTimeout(timeout*time.Second),
				//redis.DialWriteTimeout(timeout*time.Second))
			)
			if err != nil {
				return nil, err
			}
			return con, nil
		},
	}
}

func GetRedis() *redis.Pool {
	return redisClient
}
