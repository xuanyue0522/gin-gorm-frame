package config

import (
	"errors"
	"fmt"
	"gin-gorm-frame/utils/tools"
	"github.com/go-redis/redis"
)

type Redis struct {
	Alias    string `yaml:"alias"`
	Addr     string `yaml:"addr"`
	Password string `yaml:"password"`
	DbIndex  int    `yaml:"db_index"`
	MaxIdle  int    `yaml:"max_idle"`
	MaxOpen  int    `yaml:"max_open"`
}

type RedisConf struct {
	DefaultAlias string   `yaml:"default_alias"`
	Connections  []*Redis `yaml:"connections"`
}

// InitRedis 初始化redis
func InitRedis(defaultAlias string, confList []*Redis) (map[string]*redis.Client, []error) {

	// 是否含有默认的db别名
	var haveDefaultRedisAlias bool = false

	var redisMapClient = make(map[string]*redis.Client)
	var errList []error

	for _, conf := range confList {
		client := redis.NewClient(&redis.Options{
			Addr:         conf.Addr,
			Password:     conf.Password,
			DB:           conf.DbIndex,
			PoolSize:     conf.MaxOpen,
			MinIdleConns: conf.MaxIdle,
		})

		// ping
		if r, _ := client.Ping().Result(); r != "PONG" {
			errList = append(errList, errors.New(fmt.Sprintf("redis connect error: %s - db: %d", conf.Addr, conf.DbIndex)))
			continue
		}
		redisMapClient[conf.Alias] = client

		if !haveDefaultRedisAlias {
			// 判断当前数据库别名是否是default
			if conf.Alias == defaultAlias {
				haveDefaultRedisAlias = true
			}
		}
	}

	if !haveDefaultRedisAlias {
		// 处理错误
		tools.HandlePanicError(errors.New(fmt.Sprintf("There must be a redis connection with the alias %s.", defaultAlias)))
	}
	return redisMapClient, errList
}
