package store

import (
	"github.com/go-redis/redis"
	"gorm.io/gorm"
)

// 设置数据库别名
var dbAlias string = "default"

type DbBaseRepo struct {
	Db *gorm.DB
}

type RedisBaseRepo struct {
	Redis *redis.Client
}
