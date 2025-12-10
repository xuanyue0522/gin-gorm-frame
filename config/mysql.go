package config

import (
	"fmt"
	"github.com/samber/lo"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Mysql 配置
type Mysql struct {
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

// GetDsn 获取数据库连接字符串
func (m *Mysql) GetDsn() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		m.User, m.Password, m.Host, m.Port, m.Database, m.Charset)
}

// 初始化mysql
func InitMysql(conf *Mysql) (*gorm.DB, error) {
	conf.MaxIdle = lo.Max([]int{conf.MaxIdle + 1, 5})
	conf.MaxOpen = lo.Max([]int{conf.MaxOpen + 1, 10})

	dsn := conf.GetDsn()

	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		return nil, err
	}

	sqlDb, err := db.DB()
	if err != nil {
		return nil, err
	}

	if err = sqlDb.Ping(); err != nil {
		return nil, err
	}

	sqlDb.SetMaxOpenConns(conf.MaxOpen)
	sqlDb.SetMaxIdleConns(conf.MaxIdle)

	return db, nil
}
