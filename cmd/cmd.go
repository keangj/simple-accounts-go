package cmd

import (
	"log"
	"simple-accounts/internal/database"
	"simple-accounts/internal/router"
)

func RunServer() {
	database.Connect()
	database.CreateTables()
	// defer database.Close()
	r := router.New()
	// 监听 0.0.0.0:8080 端口
	err := r.Run(":8080")
	if err != nil {
		log.Fatalln(err)
	}
}
