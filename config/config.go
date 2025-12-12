package config

import (
	"flag"
	"fmt"
	"github.com/goccy/go-yaml"
	"github.com/gogf/gf/util/gconv"
	"github.com/spf13/viper"
	"os"
	"time"
)

const (
	ServerName     = "project_name"
	ServerFullName = "project_name"
)

var (
	etcdKey         = fmt.Sprintf("/configs/%s/system", ServerName)
	etcdAddr        string
	localConfigPath string
	GlobalConfig    Config
)

type Config struct {
	Server Server    `yaml:"server"`
	Db     DbConfig  `yaml:"db"`
	Redis  RedisConf `yaml:"redis"`
}

func init() {
	// 获取命令行参数
	flag.StringVar(&localConfigPath, "c", "config.yml", "default config path")
	flag.StringVar(&etcdAddr, "r", os.Getenv("ETCD_ADDR"), "DEFAULT etcd address")
}

// 初始化配置信息
func InitConfig() *Config {

	var (
		err      error
		tempConf = &Config{}
		vipConf  = viper.New()
	)

	// 解析命令行参数
	flag.Parse()

	// 如果存在etcd地址，则从etcd中获取配置
	if etcdAddr != "" {
		tempConf, err := getFromRemoteAndWatchUpdate(vipConf)
		if err != nil {
			panic(err)
		}
		return tempConf
	}

	// 从本地获取
	tempConf, err = getFromLocal()
	if err != nil {
		panic(err)
	}
	return tempConf
}

// 从etcd中获取配置，并监听配置变更
func getFromRemoteAndWatchUpdate(v *viper.Viper) (*Config, error) {
	tempConf := Config{}
	if err := v.AddRemoteProvider("etcd3", etcdAddr, etcdKey); err != nil {
		return nil, err
	}

	if err := v.ReadRemoteConfig(); err != nil {
		return nil, err
	}

	// 反序列化到结构体
	if err := v.Unmarshal(&tempConf); err != nil {
		return nil, err
	}

	// 开启一个协程监听配置变更
	go func() {
		for {
			time.Sleep(time.Minute * 1)
			if err := v.WatchRemoteConfig(); err == nil {
				_ = v.Unmarshal(&GlobalConfig)
				fmt.Println(">>> etcd config hot-reloaded: ", gconv.String(GlobalConfig)) // gconv 类型转换工具包
			}
		}
	}()
	return &tempConf, nil
}

// 从本地获取配置
func getFromLocal() (*Config, error) {
	tempConf := Config{}

	if _, err := os.Stat(localConfigPath); err == nil { // os.Stat 获取文件或目录的状态信息
		content, err := os.ReadFile(localConfigPath)
		if err != nil {
			return nil, err
		}

		// 反序列化到结构体
		if err = yaml.Unmarshal(content, &tempConf); err != nil {
			return nil, err
		}
		return &tempConf, nil
	}
	return nil, fmt.Errorf("config file not found: %s", localConfigPath)
}
