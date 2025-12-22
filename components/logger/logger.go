package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"time"
)

var SugaredLogger *zap.SugaredLogger
var logger *zap.Logger
var app string = "app-name"

func InitLogger(conf *LogConfig, appName string) error {

	// 设置全局app名称
	if appName != "" {
		app = appName
	}

	// 创建多个输出目标
	var cores []zapcore.Core

	// 编码器配置
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "message",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder, // 大写级别编码器
		EncodeTime:     customTimeEncoder,           // 自定义时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder, // 短路径编码器
	}

	// 根据配置选择编码器
	var encoder zapcore.Encoder
	if conf.Format == "json" {
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	} else {
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	}

	// 设置日志级别
	level := getZapLevel(conf.Level)

	// 输出到文件（支持日志轮转）
	if conf.LogDir != "" {
		// 确保目录存在
		if err := os.MkdirAll(conf.LogDir, os.ModePerm); err != nil {
			return err
		}

		fileWriteSyncer := getFileLogWriter(conf)
		fileCore := zapcore.NewCore(encoder, fileWriteSyncer, level)
		cores = append(cores, fileCore)
	}

	// 判断是否开启控制台输出
	if conf.ConsoleOutput {
		consoleWriteSyncer := zapcore.AddSync(os.Stdout)
		consoleCore := zapcore.NewCore(encoder, consoleWriteSyncer, level)
		cores = append(cores, consoleCore)
	}

	// 如果 cores 为空，则默认输出到控制台
	if len(cores) == 0 {
		consoleWriteSyncer := zapcore.AddSync(os.Stdout)
		consoleCore := zapcore.NewCore(encoder, consoleWriteSyncer, level)
		cores = append(cores, consoleCore)
	}

	// 创建核心
	core := zapcore.NewTee(cores...)

	// 创建logger
	logger = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	SugaredLogger = logger.Sugar()

	// 替换 zap 的全局 logger
	zap.ReplaceGlobals(logger)
	return nil
}

// 获取文件日志写入器 (支持轮转)
func getFileLogWriter(conf *LogConfig) zapcore.WriteSyncer {

	timeStr := time.Now().Format("20060102")
	filename := conf.FilenamePrefix + timeStr + ".log"

	lumberJackLogger := &lumberjack.Logger{
		Filename:   conf.LogDir + "/" + filename,
		MaxSize:    conf.MaxSize,    // MB
		MaxBackups: conf.MaxBackups, // 最大备份数量
		MaxAge:     conf.MaxAge,     // 最大保留天数
		Compress:   conf.Compress,   // 是否压缩
		LocalTime:  true,            // 使用本地时间
	}
	return zapcore.AddSync(lumberJackLogger)
}

// Sync 刷新日志缓冲区
func Sync() {
	if logger != nil {
		_ = logger.Sync()
	}
}

// getLogger 获取 logger
func getLogger() *zap.Logger {
	return zap.L()
}

func Debug(msg string, fields ...zap.Field) {
	getLogger().Debug(msg, extFields(fields)...)
}

func Info(msg string, fields ...zap.Field) {
	getLogger().Info(msg, extFields(fields)...)
}

func Warn(msg string, fields ...zap.Field) {
	getLogger().Warn(msg, extFields(fields)...)
}

func Error(msg string, fields ...zap.Field) {
	getLogger().Error(msg, extFields(fields)...)
}

// 扩展额外日志字段
func extFields(fields []zap.Field) []zap.Field {
	fields = append(fields, zap.String("app", app), zap.Int64("timestamp", time.Now().Unix()))
	return fields
}
