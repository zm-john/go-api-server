package database

import (
	"fmt"
	"github.com/jinzhu/gorm"

	"best.me/config"
	// import mysql driver
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// MysqlOrm connection
var MysqlOrm *gorm.DB

// Migrate 创建表
func Migrate(model interface{}) {
	MysqlOrm.AutoMigrate(model)
}

func init() {
	var err error
	config := config.GetMysqlConfig()
	dsn := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.Username,
		config.Password,
		config.Host,
		config.Port,
		config.Database,
	)
	MysqlOrm, err = gorm.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
}
