package store

import (
	"github.com/go-redis/redis"
	"gorm.io/gorm"
)

type BaseRepo struct {
	db    *gorm.DB
	redis *redis.Client
}
