package main

import (
	"log"
	"simple-accounts/internal/router"
)

func main() {
	r := router.New()
	// 监听 0.0.0.0:8080 端口
	err := r.Run(":8080")
	if err != nil {
		log.Fatalln(err)
	}
}
