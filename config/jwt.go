package config

// JWT config
type JWT struct {
	Key string `ini:"key"`
	TTL int64  `ini:"ttl"`
}

// GetJWTConfig 获取 JWT Config
func GetJWTConfig() JWT {
	return conf.JWT
}
