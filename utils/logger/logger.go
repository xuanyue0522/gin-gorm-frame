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
