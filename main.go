package main

import (
	"gin-gorm-frame/components/logger"
	"gin-gorm-frame/config"
	"gin-gorm-frame/server"
	"github.com/go-redis/redis"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func main() {

	// 初始化应用
	app := initApp()

	defer logger.Sync()

	// 启动http服务
	server.StartHttpServer(app.conf, app.dbMapClient, app.rdsMapClient)
}

type app struct {
	conf         *config.Config
	dbMapClient  map[string]*gorm.DB
	rdsMapClient map[string]*redis.Client
}

// 初始化应用
func initApp() *app {
	// 初始化配置信息
	conf := config.InitConfig()

	// 初始化日志
	err := logger.InitLogger(&conf.LogConf, conf.Server.AppName)
	if err != nil {
		panic(err)
	}

	// 初始化mysql
	dbMapClient, dbErrList := config.InitMysql(conf.Db.DefaultAlias, conf.Db.Connections)
	if len(dbErrList) > 0 {
		logger.Error("mysql connect error", zap.Errors("errList", dbErrList))
	}
	// 记录日志
	logger.Debug("all db(mysql) connect success")

	// 初始化redis
	rdsMapClient, redisErrList := config.InitRedis(conf.Redis.DefaultAlias, conf.Redis.Connections)
	if len(redisErrList) > 0 {
		logger.Error("redis connect error", zap.Errors("errList", redisErrList))
	}
	logger.Debug("redis connect success")

	return &app{
		conf:         conf,
		dbMapClient:  dbMapClient,
		rdsMapClient: rdsMapClient,
	}
}
