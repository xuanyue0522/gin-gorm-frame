package admin

import (
	"gin-gorm-frame/adaptor"
	"gin-gorm-frame/api/admin/v1"
	"gin-gorm-frame/api/admin/v2"
)

type Ctrl struct {
	adaptor adaptor.IAdaptor

	// v1控制器
	V1 *v1.Ctrl

	// v2控制器
	V2 *v2.Ctrl
}

func NewCtrl(adaptor adaptor.IAdaptor) *Ctrl {

	return &Ctrl{
		adaptor: adaptor,

		// 注册v1控制器
		V1: v1.RegisterCtrl(adaptor),

		// 注册v2控制器
		V2: v2.RegisterCtrl(adaptor),
	}
}
