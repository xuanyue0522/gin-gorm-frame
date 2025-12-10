package server

import (
	"context"
	"errors"
	"fmt"
	"gin-gorm-frame/adaptor"
	"gin-gorm-frame/config"
	"gin-gorm-frame/middleware"
	"gin-gorm-frame/router"
	"gin-gorm-frame/utils/logger"
	"gin-gorm-frame/utils/tools"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

type App struct {
	server *gin.Engine
	addr   string
}

func newApp(port int, router router.IRouter) *App {
	// 设置gin运行模式
	gin.SetMode(gin.ReleaseMode)

	// 创建引擎
	engine := gin.New()

	// Recover 中间件，全局捕获panic
	engine.Use(gin.Recovery())

	// 日志中间件，自定义过滤器，某些接口不需要记录日志
	engine.Use(middleware.AccessLogMiddleware(router.AccessRecordFilter))

	// 注册业务路由
	router.Register(engine)

	return &App{
		server: engine,
		addr:   ":" + strconv.Itoa(port),
	}
}

// StartHttpServer 启动http服务
func StartHttpServer(conf *config.Config, db *gorm.DB, redis *redis.Client) {

	// 创建http服务器
	app := newApp(conf.Server.HttpPort,
		router.NewRouter(
			conf,
			adaptor.NewAdaptor(conf, db, redis),
			func() error {
				err := func() error {
					pingDb, err := db.DB()
					tools.HandlePanicError(err)
					return pingDb.Ping()
				}()
				if err != nil {
					return errors.New("mysql connect failed")
				}
				return redis.Ping().Err()
			},
		),
	)

	// 启动服务
	app.run()
}

// Run 将服务跑起来
func (app *App) run() {
	srv := &http.Server{
		Addr:    app.addr,
		Handler: app.server,
	}

	// 异步启动服务
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen err: %v\n", err)
		}
	}()

	logger.Debug(fmt.Sprintf("server started, endpoint: http://localhost%s", app.addr))

	closeChan := make(chan os.Signal, 1)
	signal.Notify(closeChan, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	// 管道阻塞，等待信号
	msg := <-closeChan
	logger.Warn("server closing:", zap.String("msg", msg.String()))
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second) // 等待5秒超时
	defer cancel()

	// 关闭服务
	_ = srv.Shutdown(ctx)
}
