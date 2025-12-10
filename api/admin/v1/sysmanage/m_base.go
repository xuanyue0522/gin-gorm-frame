package sysmanage

import (
	"gin-gorm-frame/adaptor"
	"gin-gorm-frame/service/admin/sysmanage"
)

type Ctrl struct {
	adaptor adaptor.IAdaptor
	service *sysmanage.Service
}

func RegisterCtrl(adaptor adaptor.IAdaptor) *Ctrl {
	return &Ctrl{
		adaptor: adaptor,
		service: sysmanage.RegisterService(adaptor),
	}
}
