package main

import (
	"gin-gorm-frame/config"
	"gin-gorm-frame/server"
	"gin-gorm-frame/utils/logger"
	"gin-gorm-frame/utils/tools"
	"github.com/go-redis/redis"
	"gorm.io/gorm"
)

func main() {

	// 初始化应用
	app := initApp()

	// 启动http服务
	server.StartHttpServer(app.conf, app.dbMapClient, app.rdsClient)
}

type app struct {
	conf        *config.Config
	dbMapClient map[string]*gorm.DB
	rdsClient   *redis.Client
}

// 初始化应用
func initApp() *app {
	// 初始化配置信息
	conf := config.InitConfig()

	// 设置日志级别
	logger.SetLevel(conf.Server.LogLevel)

	// 初始化mysql
	dbMapClient := config.InitMysql(conf.Db)
	// 记录日志
	logger.Debug("all db(mysql) connect success")

	// 初始化redis
	rdsClient, err := config.InitRedis(&conf.Redis)
	tools.HandlePanicError(err)
	logger.Debug("redis connect success")

	return &app{
		conf:        conf,
		dbMapClient: dbMapClient,
		rdsClient:   rdsClient,
	}
}
