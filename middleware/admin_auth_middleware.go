package middleware

import (
	"errors"
	"gin-gorm-frame/api"
	"gin-gorm-frame/common"
	"gin-gorm-frame/consts"
	"github.com/gin-gonic/gin"
	"net/http"
)

type TokenAdminFun func(ctx *gin.Context, token string) (*common.AdminUser, error)

func AdminAuthMiddleware(filter func(*gin.Context) bool, getTokenFun TokenAdminFun) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if filter != nil && !filter(ctx) {
			ctx.Next()
			return
		}

		token := ctx.GetHeader(consts.AdminTokenKey)
		if len(token) == 0 {
			api.RespError(ctx, http.StatusUnauthorized, common.AuthError.WithErr(errors.New("token 不能为空")))
			ctx.Abort() // 阻止后续中间件执行
			return
		}
		user, err := getTokenFun(ctx, token)
		if err != nil {
			api.RespError(ctx, http.StatusUnauthorized, common.AuthError.WithErr(err))
			ctx.Abort()
			return
		}

		// 保存用户信息到当前上下文中
		ctx.Set(consts.AdminUserKey, user)
		ctx.Next()
	}
}
