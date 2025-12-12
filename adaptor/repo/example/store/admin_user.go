package store

import (
	"context"
	"gin-gorm-frame/adaptor"
	"gin-gorm-frame/adaptor/redis"
	"gin-gorm-frame/adaptor/repo/default/model"
	"gin-gorm-frame/adaptor/repo/default/query"
	"gin-gorm-frame/do/admin"
	"time"
)

type AdminUser struct {
	DbBaseRepo
	RedisBaseRepo
}

func NewAdminUser(adaptor adaptor.IAdaptor) *AdminUser {

	redisClient, err := adaptor.GetRedis(redis.DefaultRedisAlias)
	if err != nil {
		panic("redis-alias(default) instance is not exist")
	}

	return &AdminUser{
		DbBaseRepo: DbBaseRepo{
			Db: adaptor.GetDB(dbAlias),
		},
		RedisBaseRepo: RedisBaseRepo{
			Redis: redisClient,
		},
	}
}

type IAdminUser interface {
	CreateUser(ctx context.Context, req *admin.CreateUser) (int64, error)
}

// CreateUser 创建后台用户
func (a *AdminUser) CreateUser(ctx context.Context, req *admin.CreateUser) (int64, error) {

	timeNow := time.Now().UnixMilli() // 获取当前时间戳（毫秒）

	qs := query.Use(a.Db).AdminUser

	addObj := &model.AdminUser{
		Name:      req.Name,
		Username:  req.UserName,
		Password:  req.Password,
		Mobile:    req.Mobile,
		Status:    req.Status,
		Sex:       req.Sex,
		CreatedAt: timeNow,
		UpdatedAt: timeNow,
		CreateBy:  req.AdminUserId,
		UpdateBy:  req.AdminUserId,
	}
	err := qs.WithContext(ctx).Create(addObj)
	if err != nil {
		return 0, err
	}
	return addObj.ID, nil
}
