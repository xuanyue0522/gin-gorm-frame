package common

import (
	"gin-gorm-frame/common"
	"gin-gorm-frame/components/logger"
	"gin-gorm-frame/consts"
	"github.com/gin-gonic/gin"
)

// GetAdminUserFromCtx 从gin上下文中获取管理员用户信息
func GetAdminUserFromCtx(ctx *gin.Context) *common.AdminUser {
	user, exist := ctx.Get(consts.AdminUserKey)
	if !exist {
		return nil
	}
	adminUser, ok := user.(*common.AdminUser)
	if !ok {
		logger.Error("GetAdminUserFromCtx: 断言失败，user is not *common.AdminUser")
		return nil
	}
	return adminUser
}
