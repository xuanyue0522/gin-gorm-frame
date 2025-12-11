package defaultstore

import (
	"context"
	"gin-gorm-frame/adaptor"
	"gin-gorm-frame/adaptor/repo/model"
	"gin-gorm-frame/adaptor/repo/query"
	"gin-gorm-frame/adaptor/repo/store"
	"gin-gorm-frame/do/admindo"
	"time"
)

type AdminUser struct {
	store.DbBaseRepo
	store.RedisBaseRepo
}

func NewAdminUser(adaptor adaptor.IAdaptor) *AdminUser {

	return &AdminUser{
		DbBaseRepo: store.DbBaseRepo{
			Db: adaptor.GetDB(dbAlias),
		},
		RedisBaseRepo: store.RedisBaseRepo{
			Redis: adaptor.GetRedis(),
		},
	}
}

type IAdminUser interface {
	CreateUser(ctx context.Context, req *admindo.CreateUser) (int64, error)
}

// CreateUser 创建后台用户
func (a *AdminUser) CreateUser(ctx context.Context, req *admindo.CreateUser) (int64, error) {

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
