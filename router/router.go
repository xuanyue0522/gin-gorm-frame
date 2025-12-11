package router

import (
	"gin-gorm-frame/adaptor"
	"gin-gorm-frame/api/admin"
	"gin-gorm-frame/config"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type IRouter interface {
	Register(engine *gin.Engine)
	SpanFilter(r *gin.Context) bool
	AccessRecordFilter(r *gin.Context) bool
}

type Router struct {
	FullPPROF bool
	rootPath  string
	conf      *config.Config
	checkFunc func() error

	// 管理后台控制器
	admin *admin.Ctrl
}

func NewRouter(conf *config.Config, adaptor adaptor.IAdaptor, checkFunc func() error) *Router {
	return &Router{
		FullPPROF: conf.Server.EnablePprof,
		rootPath:  "/api", // 路由根路径
		conf:      conf,
		checkFunc: checkFunc,

		// 注册管理后台控制器
		admin: admin.NewCtrl(adaptor), // 管理后台控制器
	}
}

// 注册业务路由组 (每新增一个业务模块，就新增一个路由组)
func (r *Router) registerServiceRouteGroup(root *gin.RouterGroup) {

	// 注册管理后台路由组
	RegisterAdminRouteGroup(r, root)
}

// 检测服务器是否正常运行
func (r *Router) checkServer() func(*gin.Context) {
	return func(ctx *gin.Context) {
		err := r.checkFunc()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	}
}

// Register 注册路由
func (r *Router) Register(app *gin.Engine) {
	if r.conf.Server.EnablePprof { // 说明开启了pprof （go内置的性能分析工具）

		// 启动性能分析工具
		SetupPprof(app, "/debug/pprof")
	}
	app.Any("/ping", r.checkServer())

	root := app.Group(r.rootPath)

	// 注册业务路由组
	r.registerServiceRouteGroup(root)
}

// SpanFilter 过滤器
func (r *Router) SpanFilter(ctx *gin.Context) bool {
	path := strings.Replace(ctx.Request.URL.Path, r.rootPath, "", 1)
	_, ok := AdminAuthWhiteList[path]
	return !ok
}

func (r *Router) AccessRecordFilter(ctx *gin.Context) bool {
	return true
}
