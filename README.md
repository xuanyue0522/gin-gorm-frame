## gin-gorm-frame

[![Go Version](https://img.shields.io/badge/Go-1.24.0-blue.svg)](https://golang.org)
[![Gin Framework](https://img.shields.io/badge/Gin-v1.11.0-green.svg)](https://github.com/gin-gonic/gin)
[![GORM](https://img.shields.io/badge/Gorm-v1.31.1-orange.svg)](https://gorm.io)

一个基于 Gin + gorm + redis 的企业级 Go Web 应用模板，采用清晰地分层架构设计，内置依赖注入、健康检查、性能分析、配置管理等企业级功能。

## ✨ 特性

- 🏗️ **清晰的分层架构**：api -> dto → service → do → store -> model
- 🔧 **依赖注入**
- 🌐 **多数据库实例支持**：支持配置连接多个mysql
- 📦 **Redis 缓存**：支持配置连接多个 redis 
- 📊 **健康检查**：完整的健康检查机制
- 🔧 **配置管理**：基于 Viper 的配置管理
- 📝 **结构化日志**：基于 Zap 的高性能日志
- 🚀 **优雅启停**：支持优雅关闭
- 🔄 **自动生成代码**：利用 Gen Tool 自动生成 model 和 query 代码

## 快速开始

### 1. 克隆项目

```bash
# 克隆项目
git clone https://github.com/xuanyue0522/gin-gorm-frame.git gin-gorm-frame
# 进入项目
cd gin-gorm-frame
```

### 2. 配置环境

```bash
# 拷贝示例配置文件
cp config.example.yml config.yml
# 编辑配置文件
vim config.yml
```

### 3. 安装依赖

```bash
go mod tidy
```

### 4. 修改 gen.yml
adaptor/repo/default/gen.yaml
根据自己实际情况修改

### 5. 自动生成代码

```bash
make -f Makefile
```

### 5. 启动项目

```bash
go run main.go
```

## 目录结构
```
gin-gorm-frame/
├── main.go                           # 程序入口
├── go.mod                            # Go模块定义
├── go.sum                            # 依赖版本锁定
├── README.md                         # 项目说明文档
├── config.example.yml                # 配置文件示例
├── config.yml                        # 配置文件
├── Makefile                          # 构建脚本
├── package.json                      # npm包管理文件
├── package-lock.json                 # npm包管理锁文件
├── .gitignore                        # git 忽略文件
├── .cz-config.js                     # git 提交规范配置文件
├── adaptor/                          # 适配器
│   ├── redis/                        # redis 相关操作包
│   │   ├── redis.go                  # redis 操作必要配置
│   │   └── example_redis.go          # redis 操作示例文件
│   ├── repo/                         # 数据库操作相关
│   │   ├── default/                  # 默认数据库连接-数据库操作
│   │   │   ├──model/                 # 模型（gen-tool 自动生成）
│   │   │   ├──query/                 # 模型（gen-tool 自动生成
│   │   │   ├──store/                 # 仓储包（自己的操作数据库的逻辑）
│   │   │   │    └── base.go          # 仓储包基础配置 
│   │   │   └── gen.yaml              # gen-tool 配置文件
│   │   │── example/                  # 示例数据库连接-数据库操作
│   │   │   ├──model/                 # 模型（gen-tool 自动生成）
│   │   │   ├──query/                 # 模型（gen-tool 自动生成
│   │   │   ├──store/                 # 仓储包（自己的操作数据库的逻辑）
│   │   │   │    └── base.go          # 仓储包基础配置 
│   │   │   └── gen.yaml              # gen-tool 配置文件
│   │   └── rpc/                      # rpc 相关操作
│   └── adaptor.go                    # 适配器文件
├── api/                              # 接口包
│   ├── admin/                        # 后台管理模块
│   │   ├── common/                   # 后台模块公共部分
│   │   │   └── common.go             # 后台模块公共文件
│   │   ├── v1/                       # 后台模块 v1 版本控制器
│   │   │   ├── sysmanage/            # 后台系统管理模块控制器包
│   │   │   │   ├── m_base.go         # 控制器功能模块(m)基础文件
│   │   │   │   └── user_ctrl.go      # 后台用户管理控制器
│   │   │   └── v_base.go             # 控制器版本(v)基础文件
│   │   ├── v2/                       # 后台模块 v2 版本控制器
│   │   │   ├── sysmanage/            # 后台系统管理模块控制器包
│   │   │   │   ├── m_base.go         # 控制器功能模块(m)基础文件
│   │   │   │   └── user_ctrl.go      # 后台用户管理控制器
│   │   │   └── v_base.go             # 控制器版本(v)基础文件
│   │   └── p_base.go                 # 控制器项目(p)基础文件
│   ├── web/                          # 前台模块（文件结构将与 admin 模块一致）
│   └── response.go                   # 统一响应封装
│    
├── common/                           # 公共包
│   └── errors.go                     # 错误封装与错误码定义
├── config/                           # 配置包
│   ├── config.go                     # 主配置文件
│   ├── redis.go                      # redis 配置文件
│   ├── server.go                     # 服务器配置文件
│   └── mysql.go                      # mysql配置文件
├── consts/                           # 常量包
│   └── consts.go                     # 常量配置
├── do/                               # 数据对象
│   ├── admin/                        # 后台模块数据对象
│   │   └── admin_user_do.go          # 后台用户数据对象
│   ├── web/                          # 前台模块数据对象
│   └── readme.md                     # 数据对象说明文件
├── dto/                              # 数据传输对象
│   ├── admin/                        # 后台模块数据传输对象
│   │   └── admin_user_dto.go         # 后台用户数据传输对象
│   ├── web/                          # 前台模块数据传输对象
│   └── readme.md                     # 数据传输对象说明文件
├── middleware/                       # 中间件
│   ├── access_middleware.go          # 运行日志中间件
│   └── admin_auth_middleware.go      # 管理后台用户权限中间件
├── router/                           # 路由
│   ├── api_admin.go                  # 管理后台路由
│   ├── api_web.go                    # web 前台路由
│   ├── pprof.go                      # 性能分析工具路由
│   ├── route.go                      # 核心主路由
│   └── white_list.go                 # 路由白名单
├── server/                           # 服务
│   └── http.go                       # http 服务
├── service/                          
│   ├── admin/                        # 后台管理模块 service
│   │   └── sysmanage/                # 后台系统管理模块 service 包
│   │       ├── m_service.go          # service 模块(m)基础文件
│   │       └── user_service.go       # 后台用户管理 service
│   │  
│   └── web/                          # 前台模块 service
├── sql/                          
│   └── db.sql                        # 数据库示例sql文件
└── utils/                            # 工具箱
    ├── captcha/                      # 滑块验证码
    │   └── captcha.go                
    ├── logger/                       # 日志
    │   └── logger.go                 
    ├── pool/                         # 协程池
    │   └── pool.go                   
    └── tools/                        # 工具
        └── error.go                  # 异常判断
        └── tool.go                   # 工具
```

## 🛠️ 技术栈

| 技术 | 版本 | 描述 |
|------|------|------|
| **Go** | 1.24 | 编程语言 |
| **Gin** | 1.11.0 | Web框架 |
| **GORM** | 1.31.1 | ORM框架 |
| **MySQL/PostgreSQL** | - | 关系型数据库 |
| **Redis** | - | 缓存数据库 |
| **Viper** | 1.21.0 | 配置管理 |
| **Zap** | 1.27.1 | 结构化日志 |

### 分层架构

```
┌─────────────────────────────────┐
│          HTTP Layer             │
│     (Gin Router & Middleware)   │
├─────────────────────────────────┤
│        API - Controller Layer   │
│    (Request/Response Handling)  │
├─────────────────────────────────┤
│           DTO Layer             │
│    (Data - teansfer - Object)   │
├─────────────────────────────────┤
│          Service Layer          │
│        (Business Logic)         │
├─────────────────────────────────┤
│          DO Layer               │
│       (Data Objects)            │
├─────────────────────────────────┤
│          Store Layer            │
│         (Data Logic)            │
├─────────────────────────────────┤
│         Model Layer             │
│    (Database Models)            │
└─────────────────────────────────┘
```
## 响应示例
```json
{
  "code": 200,
  "msg": "OK",
  "err_msg": "",
  "data": {
    "id": 1
  }
}
```

## 配置说明
config.yml
```yaml
# http 服务
server:
  http_port: 8081
  enable_pprof: true
  log_level: debug
  env: dev

# db 数据库
db:
  default_alias: default # 默认数据库别名
  connections:           # mysql 连接
    - db1:               # mysql库1 连接
      alias: default
      dialect: mysql
      user: root
      password: 123456
      host: 127.0.0.1
      port: 3306
      database: gin-gorm-frame
      charset: utf8mb4
      show_mysql: true
      max_open: 20
      max_idle: 5

    - db2:              # mysql库2 连接
      alias: db2
      dialect: mysql
      user: root
      password: 123456
      host: 127.0.0.1
      port: 3306
      database: gin-gorm-frame
      charset: utf8mb4
      show_mysql: true
      max_open: 20
      max_idle: 5

# redis
redis:
  default_alias: default  # 默认 redis 连接别名
  connections:            # redis 连接
    - redis1:             # redis库1 连接
      alias: default
      addr: 127.0.0.1:6479
      password: ""
      db_index: 15
      max_idle: 2
      max_open: 10

    - redis2:             # redis库2 连接
      alias: session
      addr: 127.0.0.1:6479
      password: ""
      db_index: 14
      max_idle: 2
      max_open: 10

```
## 🔧 开发指南

### 编码规范
- 使用 `gofmt` 格式化代码
- 遵循 Go 官方命名规范
- 优先使用小接口，遵循单一职责原则

## 错误处理
项目使用统一的错误码和响应格式：

```go
package common

type Errno struct {
	Code   int    // 错误码
	Msg    string // 错误提示（用于在前端展示）
	ErrMsg string // 详细错误信息或原因（用于调试）
}

func (e Errno) Error() string {
	return e.Msg
}

// WithMsg 对简短错误信息进行补充说明
func (e Errno) WithMsg(msg string) Errno {
	e.Msg = e.Msg + " " + msg
	return e
}

// WithErr 错误的详细错误信息
func (e Errno) WithErr(rawErr error) Errno {
	var msg string
	if rawErr != nil {
		msg = rawErr.Error()
	}
	e.ErrMsg = msg
	return e
}

func (e Errno) IsOk() bool {
	return e.Code == 200
}

var (

	/*
		服务器级别错误
	*/
	OK              = Errno{Code: 200, Msg: "OK"}
	ServerError     = Errno{Code: 500, Msg: "Server Error"}
	ParamError      = Errno{Code: 400, Msg: "Param Error"}
	AuthError       = Errno{Code: 401, Msg: "Auth Error"}
	PermissionError = Errno{Code: 403, Msg: "Permission Error"}

	/*
		数据库级别错误
		错误码从10000开始
	*/
	DatabaseError = Errno{Code: 10000, Msg: "Database Error"}
	RedisError    = Errno{Code: 10001, Msg: "Redis Error"}

	/*
		工具级别错误
		错误码从20000开始
	*/
	InvalidCaptchaError = Errno{Code: 20000, Msg: "滑块校验失败，请刷新重试"}

	/*
		用户级别错误
		错误码从30000开始
	*/
	UserNotFound      = Errno{Code: 30000, Msg: "User Not Found"}
	UserAlreadyExists = Errno{Code: 30001, Msg: "User Already Exists"}
	UserPasswordError = Errno{Code: 30002, Msg: "User Password Error"}
	UserNotFoundError = Errno{Code: 30003, Msg: "User Not Found"}
)

```

## 日志
```go
package logger

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.Logger

// 创建原子级别的便利函数
var atom = zap.NewAtomicLevelAt(zap.DebugLevel)

func init() {
	config := zap.Config{
		Level:       atom,
		Development: false,  // 开发者模式
		Encoding:    "json", // 指定 JSON 编码
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey: "msg",
			LevelKey:   "level",
			TimeKey:    "time",
			CallerKey:  "caller",
			// StacktraceKey: "stacktrace",
			EncodeTime:   zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05"),
			EncodeLevel:  zapcore.LowercaseLevelEncoder,
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
	}

	tempLogger, err := config.Build()
	if err != nil {
		panic(err)
	}
	logger = tempLogger.WithOptions(zap.AddCaller(), zap.AddCallerSkip(1), zap.AddStacktrace(zap.ErrorLevel))
}

// SetLevel 设置日志级别
func SetLevel(level string) {
	// 解析日志级别
	tLevel, err := zapcore.ParseLevel(level)
	if err != nil {
		fmt.Printf("日志级别（%s）解析失败: %v\n", level, err)
		return
	}
	atom.SetLevel(tLevel)
}

func Debug(msg string, fields ...zap.Field) {
	logger.Debug(msg, fields...)
}

func Info(msg string, fields ...zap.Field) {
	logger.Info(msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
	logger.Warn(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	logger.Error(msg, fields...)
}

```
使用：
```go
package main

import (
	"gin-gorm-frame/components/logger"
	"go.uber.org/zap"
	"gin-gorm-frame/common"
)


func main()  {
	// 记录日志
	logger.Debug("all db(mysql) connect success")
    errs := []error {
		common.ParamError,
		common.AuthError,
    }
	logger.Error("redis connect error", zap.Errors("errList", errs))   
}
```

## git 规范

安装 node 依赖：
```bash
npm install
```
使用 git cz 代替传统的 git commit,如：
```bash
git cz
? 选择本次提交类型： (Use arrow keys)
❯ feat:     添加新功能 
  fix:      修复Bug 
  docs:     文档变更 
  style:    代码格式（不影响代码运行的变动） 
  refactor: 代码重构（既不是新增功能，也不是Bug修复） 
  perf:     性能优化 
  test:     添加缺失的测试或修改现有的测试 
```

## 鸣谢

本项目设计灵感来主要自于一下项目（排名不分先后）：
* [yuanqinguo](https://github.com/yuanqinguo/flyfei/tree/master/src/golang/project/backend/mall)
* [muleiwu](https://github.com/muleiwu/go-web)