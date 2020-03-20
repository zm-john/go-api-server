package cmd

import (
	"best.me/config"
	"best.me/database"
	"best.me/middleware"
	"best.me/routes"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"io"
	"os"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "start app serve",
	Run: func(cmd *cobra.Command, args []string) {
		serve()
	},
}

func serve() {
	conf := config.GetAppConfig()
	if conf.Mode == "release" {
		// 设置生产模式
		gin.SetMode(gin.ReleaseMode)
		// 记录日志到文件
		f, _ := os.Create("run.log")
		gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	}

	r := gin.Default()
	r.Use(middleware.RecoveryWithWriter(gin.DefaultWriter))

	// 添加路由
	routes.AddRoutes(r)

	// 监听并启动服务
	r.Run(":8000")

	// 退出是关闭数据库连接
	database.MysqlOrm.Close()
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
