package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
		log.Println("setupRouter")

	})
	return r
}

func main() {
	r := setupRouter()
	// 监听 0.0.0.0:8080 端口
	r.Run(":8080")
}
