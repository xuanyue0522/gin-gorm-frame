package config

// Server http 服务配置
type Server struct {
	AppName     string `yaml:"app_name"`
	HttpPort    int    `yaml:"http_port"`
	EnablePprof bool   `yaml:"enable_pprof"`
	Env         string `yaml:"env"`
}
