package config

// App config
type App struct {
	Mode string `ini:"mode"`
}

// GetAppConfig 获取 APP Config
func GetAppConfig() App {
	return conf.App
}
