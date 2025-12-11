package config

import (
	"errors"
	"fmt"
	"gin-gorm-frame/utils/logger"
	"gin-gorm-frame/utils/tools"
	"github.com/samber/lo"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 默认的db别名
var defaultDbAlias = "default"

// DbItemConfig 配置
type DbItemConfig struct {
	Alias     string `yaml:"alias"`
	Dialect   string `yaml:"dialect"`
	User      string `yaml:"user"`
	Password  string `yaml:"password"`
	Host      string `yaml:"host"`
	Port      int    `yaml:"port"`
	Database  string `yaml:"database"`
	Charset   string `yaml:"charset"`
	ShowMysql bool   `yaml:"show_mysql"`
	MaxOpen   int    `yaml:"max_open"`
	MaxIdle   int    `yaml:"max_idle"`
}

// getDsn 获取数据库连接字符串
func getDsn(dbItem *DbItemConfig) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		dbItem.User, dbItem.Password, dbItem.Host, dbItem.Port, dbItem.Database, dbItem.Charset)
}

// InitMysql 初始化mysql
func InitMysql(dbItemList []*DbItemConfig) map[string]*gorm.DB {

	// 是否含有默认的db别名
	var haveDefaultDbAlias bool = false

	dbMapClient := make(map[string]*gorm.DB)

	for _, conf := range dbItemList {

		conf.MaxIdle = lo.Max([]int{conf.MaxIdle + 1, 5})
		conf.MaxOpen = lo.Max([]int{conf.MaxOpen + 1, 10})

		dsn := getDsn(conf)

		db, err := gorm.Open(mysql.Open(dsn))
		if err != nil {
			// 处理错误
			tools.HandlePanicError(err)
			return nil
		}

		sqlDb, err := db.DB()
		if err != nil {
			// 处理错误
			tools.HandlePanicError(err)
			return nil
		}

		if err = sqlDb.Ping(); err != nil {
			// 处理错误
			tools.HandlePanicError(err)
			return nil
		}

		sqlDb.SetMaxOpenConns(conf.MaxOpen)
		sqlDb.SetMaxIdleConns(conf.MaxIdle)

		// 记录日志
		logger.Debug(fmt.Sprintf("db (%s - %s) connect success", conf.Alias, conf.Database))

		// 记录数据库连接到map中
		dbMapClient[conf.Alias] = db

		if !haveDefaultDbAlias {
			// 判断当前数据库别名是否是default
			if conf.Alias == defaultDbAlias {
				haveDefaultDbAlias = true
			}
		}
	}

	if !haveDefaultDbAlias {
		// 处理错误
		tools.HandlePanicError(errors.New(fmt.Sprintf("There must be a database with the alias %s.", defaultDbAlias)))
	}
	return dbMapClient
}
