package sysmanage

import (
	"gin-gorm-frame/adaptor"
	"gin-gorm-frame/adaptor/repo/store/defaultstore"
)

type Service struct {

	// 后台用户
	adminUser defaultstore.IAdminUser
}

func RegisterService(adaptor adaptor.IAdaptor) *Service {

	return &Service{

		// 注册后台用户repo
		adminUser: defaultstore.NewAdminUser(adaptor),
	}
}
