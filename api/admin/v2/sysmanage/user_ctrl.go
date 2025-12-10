package sysmanage

import (
	"gin-gorm-frame/api"
	"gin-gorm-frame/common"
	"gin-gorm-frame/dto/admindto"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (c *Ctrl) CreateUser(ctx *gin.Context) {
	// 获取并绑定请求数据
	req := &admindto.CreateUserReq{}
	if err := ctx.ShouldBindJSON(req); err != nil {
		api.RespError(ctx, http.StatusBadRequest, common.ParamError.WithErr(err))
		return
	}

	// 调用创建用户服务
}
