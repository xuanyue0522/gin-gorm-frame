package sysmanage

import (
	"gin-gorm-frame/adaptor"
	"gin-gorm-frame/adaptor/repo/store"
)

type Service struct {

	// 后台用户
	adminUser store.IAdminUser
}

func RegisterService(adaptor adaptor.IAdaptor) *Service {

	return &Service{

		// 注册后台用户repo
		adminUser: store.NewAdminUser(adaptor),
	}
}
