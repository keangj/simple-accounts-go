package cmd

import (
	"fmt"
	"log"
	"os"
	"simple-accounts/internal/database"
	"simple-accounts/internal/router"
	"time"

	"github.com/spf13/cobra"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "admin"
	password = "123456"
	dbname   = "simple_accounts_dev"
)

type User struct {
	ID        uint
	Email     *string
	CreatedAt time.Time
	UpdatedAt time.Time
}
type Items struct {
	ID        uint
	TagId     uint
	Amount    *float64
	CreatedAt time.Time
	UpdatedAt time.Time
}

func Run() {
	rootCmd := &cobra.Command{
		Use:   "simple-accounts",
		Short: "simple-accounts App Cli",
	}
	srvCmd := &cobra.Command{
		Use:   "server",
		Short: "Run server",
		Run: func(cmd *cobra.Command, args []string) {
			RunServer()
		},
	}
	dbCmd := &cobra.Command{
		Use:   "db",
		Short: "Run db",
	}
	lib1Cmd := &cobra.Command{
		Use:   "sql",
		Short: "use database/sql",
		Run: func(cmd *cobra.Command, args []string) {
			database.Connect()
			database.CreateTables()
			defer database.Close()
		},
	}
	lib2Cmd := &cobra.Command{
		Use:   "gorm",
		Short: "use gorm",
		Run: func(cmd *cobra.Command, args []string) {
			connectStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
			db, err := gorm.Open(postgres.Open(connectStr))
			if err != nil {
				log.Fatalln(err)
			}
			// db.Migrator().CreateTable(&User{}, &Items{}) // 创建表
			db.AutoMigrate(&User{}, &Items{}) // 自动迁移
		},
	}
	dbCmd.AddCommand(lib1Cmd)
	dbCmd.AddCommand(lib2Cmd)
	rootCmd.AddCommand(dbCmd)
	rootCmd.AddCommand(srvCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func RunServer() {
	// database.Connect()
	// database.MysqlConnect()
	// database.CreateTables()
	// database.MysqlCreateTables()
	// defer database.Close()
	// defer database.MysqlClose()
	r := router.New()
	// 监听 0.0.0.0:8080 端口
	err := r.Run(":8080")
	if err != nil {
		log.Fatalln(err)
	}
}
