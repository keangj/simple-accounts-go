package cmd

import (
	"fmt"
	"log"
	"os"
	"simple-accounts/internal/database"
	"simple-accounts/internal/email"
	"simple-accounts/internal/router"
	"time"

	"github.com/spf13/cobra"
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
	createCmd := &cobra.Command{
		Use: "create",
		Run: func(cmd *cobra.Command, args []string) {
			database.CreateTables()
		},
	}
	createMigrationsCmd := &cobra.Command{
		Use: "create:migrations",
		Run: func(cmd *cobra.Command, args []string) {
			database.CreateMigrate(args[0])
		},
	}
	mgrtCmd := &cobra.Command{
		Use: "migrate",
		Run: func(cmd *cobra.Command, args []string) {
			database.Migrate()
		},
	}
	mgrtDownCmd := &cobra.Command{
		Use: "migrate:down",
		Run: func(cmd *cobra.Command, args []string) {
			database.MigrateDown()
		},
	}
	crudCmd := &cobra.Command{
		Use: "crud",
		Run: func(cmd *cobra.Command, args []string) {
			database.Crud()
		},
	}

	emailCmd := &cobra.Command{
		Use: "email",
		Run: func(cmd *cobra.Command, args []string) {
			email.Send()
		},
	}

	database.Connect()
	defer database.Close()
	rootCmd.AddCommand(dbCmd, srvCmd, emailCmd)
	dbCmd.AddCommand(createCmd, createMigrationsCmd, mgrtCmd, mgrtDownCmd, crudCmd)

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
