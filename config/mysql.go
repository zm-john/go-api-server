package config

// Mysql Config
type Mysql struct {
	Host     string `ini:"host"`
	Port     string `ini:"port"`
	Username string `ini:"username"`
	Password string `ini:"password"`
	Database string `ini:"database"`
}

// GetMysqlConfig 获取 Mysql Config
func GetMysqlConfig() Mysql {
	return conf.Mysql
}
