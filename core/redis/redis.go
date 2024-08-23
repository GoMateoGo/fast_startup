package myrd

import (
	"context"

	"github.com/redis/go-redis/v9"
)

var MyRd *Redis

type Redis struct {
	Client *redis.Client
}

// 定义Redis结构体，包含Host、Pwd、Db三个字段
type RedisCfg struct {
	Host string // 主机地址
	Pwd  string // 密码
	Db   int    // 数据库编号
}

func NewRedis(rd *RedisCfg) {
	client := redis.NewClient(&redis.Options{
		Addr:     rd.Host,
		Password: rd.Pwd,
		DB:       rd.Db,
	})
	if _, err := client.Ping(context.Background()).Result(); err != nil {
		panic(err)
	}

	MyRd = &Redis{
		Client: client,
	}
}
