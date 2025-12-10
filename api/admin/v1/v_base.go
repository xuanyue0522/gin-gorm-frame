package v1

import (
	"gin-gorm-frame/adaptor"
	"gin-gorm-frame/api/admin/v1/sysmanage"
)

type Ctrl struct {
	adaptor adaptor.IAdaptor

	// 系统管理模块控制器
	Sysmanage *sysmanage.Ctrl
}

func RegisterCtrl(adaptor adaptor.IAdaptor) *Ctrl {

	return &Ctrl{

		// 注入adaptor（适配器）
		adaptor: adaptor,

		// 注册系统管理模块控制器
		Sysmanage: sysmanage.RegisterCtrl(adaptor),
	}
}
