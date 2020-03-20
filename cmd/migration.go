package cmd

import (
	"best.me/database"
	"best.me/models"
	"fmt"
	"github.com/spf13/cobra"
)

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "migrate database",
	Run: func(cmd *cobra.Command, args []string) {
		migrate()
	},
}

func migrate() {
	// 创建 users 表
	database.MysqlOrm.AutoMigrate(models.User{})
	database.MysqlOrm.Close()
	fmt.Println("迁移完成")
}

func init() {
	rootCmd.AddCommand(migrateCmd)
}
