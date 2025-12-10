package config

import (
	"errors"
	"github.com/go-redis/redis"
)

type Redis struct {
	Addr     string `yaml:"addr"`
	Password string `yaml:"password"`
	DbIndex  int    `yaml:"db_index"`
	MaxIdle  int    `yaml:"max_idle"`
	MaxOpen  int    `yaml:"max_open"`
}

// InitRedis 初始化redis
func InitRedis(conf *Redis) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:         conf.Addr,
		Password:     conf.Password,
		DB:           conf.DbIndex,
		PoolSize:     conf.MaxOpen,
		MinIdleConns: conf.MaxIdle,
	})

	// ping
	if r, _ := client.Ping().Result(); r != "PONG" {
		return nil, errors.New("redis connect failed")
	}
	return client, nil
}
