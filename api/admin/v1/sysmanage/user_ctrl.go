package sysmanage

import (
	"fmt"
	"gin-gorm-frame/api"
	commonCtrl "gin-gorm-frame/api/admin/common"
	"gin-gorm-frame/common"
	"gin-gorm-frame/dto/admindto"
	"github.com/gin-gonic/gin"
)

func (c *Ctrl) CreateUser(ctx *gin.Context) {

	// 获取当前管理员用户信息
	adminUser := commonCtrl.GetAdminUserFromCtx(ctx)
	if adminUser == nil {
		api.WriteResp(ctx, nil, common.AuthError.WithErr(fmt.Errorf("尚未登录或登录已过期")))
		return
	}

	// 获取并绑定请求数据
	req := &admindto.CreateUserReq{}
	if err := ctx.ShouldBindJSON(req); err != nil {
		api.WriteResp(ctx, nil, common.ParamError.WithErr(err))
		return
	}

	// 调用创建用户服务
	userId, err := c.service.CreateUser(ctx, adminUser, req)
	api.WriteResp(ctx, map[string]int64{
		"id": userId,
	}, err)
}
