package sysmanage

import (
	"context"
	"gin-gorm-frame/common"
	admindo "gin-gorm-frame/do/admin"
	admindto "gin-gorm-frame/dto/admin"
	"gin-gorm-frame/utils/logger"
	"go.uber.org/zap"
)

// CreateUser 创建后台用户
func (s *Service) CreateUser(ctx context.Context, adminUser *common.AdminUser, req *admindto.CreateUserReq) (int64, common.Errno) {

	userId, err := s.adminUser.CreateUser(ctx, &admindo.CreateUser{
		AdminUserId: adminUser.UserId,
		Name:        req.Name,
		UserName:    req.UserName,
		Password:    "",
		Mobile:      req.Mobile,
		Status:      req.Status,
		Sex:         req.Sex,
	})

	if err != nil {
		logger.Error("CreateUser error", zap.Error(err), zap.Any("req", req))
		return 0, common.DatabaseError.WithErr(err)
	}
	return userId, common.OK
}

// CreateUserV2 创建后台用户v2版本
func (s *Service) CreateUserV2(ctx context.Context, adminUser *common.AdminUser, req *admindto.CreateUserReq) (int64, error) {

	userId, err := s.adminUser.CreateUser(ctx, &admindo.CreateUser{
		AdminUserId: adminUser.UserId,
		Name:        req.Name,
		UserName:    req.UserName,
		Password:    "",
		Mobile:      req.Mobile,
		Status:      req.Status,
		Sex:         req.Sex,
	})

	if err != nil {
		logger.Error("CreateUser error", zap.Error(err), zap.Any("req", req))
		return 0, common.DatabaseError.WithErr(err)
	}
	return userId, nil
}
