package config

// AppConfig 应用程序配置
type ServerConfig struct {
	Mode string `ini:"mode"`
	Port int  `ini:"port"`
}
