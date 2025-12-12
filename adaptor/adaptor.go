package adaptor

import (
	"errors"
	"fmt"
	rAlias "gin-gorm-frame/adaptor/redis"
	"gin-gorm-frame/config"
	"github.com/go-redis/redis"
	"gorm.io/gorm"
)

// IAdaptor 适配器接口
type IAdaptor interface {
	GetConfig() *config.Config
	GetDB(dbAlias string) *gorm.DB
	GetRedis(redisAlias rAlias.RedisAlias) (*redis.Client, error)
}

type Adaptor struct {
	conf     *config.Config
	dbMap    map[string]*gorm.DB
	redisMap map[string]*redis.Client
}

// NewAdaptor 初始化adaptor（适配器）
func NewAdaptor(conf *config.Config, dbMap map[string]*gorm.DB, redisMap map[string]*redis.Client) *Adaptor {
	return &Adaptor{
		conf:     conf,
		dbMap:    dbMap,
		redisMap: redisMap,
	}
}

// GetConfig 获取全局配置
func (a *Adaptor) GetConfig() *config.Config {
	return a.conf
}

// GetDB 获取数据库连接
func (a *Adaptor) GetDB(dbAlias string) *gorm.DB {

	db, ok := a.dbMap[dbAlias]
	if !ok {
		panic(fmt.Sprintf("db-alias(%s) instance is not exist，Please check whether the alias of database configuration is consistent with the alias set in repo!", dbAlias))
	}
	return db
}

// GetRedis 获取redis连接
func (a *Adaptor) GetRedis(redisAlias rAlias.RedisAlias) (*redis.Client, error) {
	redisClient, ok := a.redisMap[string(redisAlias)]
	if !ok {
		return nil, errors.New(fmt.Sprintf("redis-alias(%s) instance is not exist，Please check whether the alias of redis configuration is consistent with the alias set in repo!", redisAlias))
	}
	return redisClient, nil
}
