package config

// Server http 服务配置
type Server struct {
	HttpPort    int    `yaml:"http_port"`
	EnablePprof bool   `yaml:"enable_pprof"`
	LogLevel    string `yaml:"log_level"`
	Env         string `yaml:"env"`
}
