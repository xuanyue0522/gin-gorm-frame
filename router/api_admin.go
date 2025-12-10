package router

import (
	"gin-gorm-frame/common"
	"gin-gorm-frame/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterAdminRouteGroup(r *Router, root *gin.RouterGroup) {

	// 后台中间件
	adminRoot := root.Group("/admin", middleware.AdminAuthMiddleware(r.SpanFilter, func(ctx *gin.Context, token string) (*common.AdminUser, error) {
		return &common.AdminUser{
			UserId: 1,
			Name:   "admin",
		}, nil
	}))

	// ========================= v1 版本 =========================

	// 后台v1版本路由组
	v1RootGroup := adminRoot.Group("/v1")
	// 后台v1版本控制器
	v1Ctrl := r.admin.V1

	/**
	后台管理员相关
	*/
	// 创建后台管理员
	v1RootGroup.POST("/user/create", v1Ctrl.Sysmanage.CreateUser)

	// ========================= v2 版本 =========================
	// 后台v2版本路由组
	v2RootGroup := adminRoot.Group("/v2")
	// 后台v2版本控制器
	v2Ctrl := r.admin.V2

	/**
	后台管理员相关
	*/
	// 创建后台管理员
	v2RootGroup.POST("/user/create", v2Ctrl.Sysmanage.CreateUser)

}
