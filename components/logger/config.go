package logger

type LogConfig struct {
	Level          string `yaml:"level"`           // 日志级别: debug, info, warn, error
	Format         string `yaml:"format"`          // 输出格式: console, json
	LogDir         string `yaml:"log_dir"`         // 日志目录
	FilenamePrefix string `yaml:"filename_prefix"` // 日志文件名
	MaxSize        int    `yaml:"max_size"`        // 单个文件最大大小(MB)
	MaxBackups     int    `yaml:"max_backups"`     // 保留旧文件的最大个数
	MaxAge         int    `yaml:"max_age"`         // 保留旧文件的最大天数
	Compress       bool   `yaml:"compress"`        // 是否压缩旧文件
	ConsoleOutput  bool   `yaml:"console_output"`  // 是否输出到控制台
}
