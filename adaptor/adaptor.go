package adaptor

import (
	"fmt"
	"gin-gorm-frame/config"
	"github.com/go-redis/redis"
	"gorm.io/gorm"
)

// IAdaptor 适配器接口
type IAdaptor interface {
	GetConfig() *config.Config
	GetDB(dbAlias string) *gorm.DB
	GetRedis() *redis.Client
}

type Adaptor struct {
	conf  *config.Config
	dbMap map[string]*gorm.DB
	redis *redis.Client
}

// NewAdaptor 初始化adaptor（适配器）
func NewAdaptor(conf *config.Config, dbMap map[string]*gorm.DB, redis *redis.Client) *Adaptor {
	return &Adaptor{
		conf:  conf,
		dbMap: dbMap,
		redis: redis,
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
func (a *Adaptor) GetRedis() *redis.Client {
	return a.redis
}
