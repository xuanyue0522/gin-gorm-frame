package store

import (
	"github.com/go-redis/redis"
	"gorm.io/gorm"
)

type DbBaseRepo struct {
	Db *gorm.DB
}

type RedisBaseRepo struct {
	Redis *redis.Client
}
