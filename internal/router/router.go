package router

import (
	"simple-accounts/config"
	"simple-accounts/internal/controller"
	"simple-accounts/internal/database"

	"simple-accounts/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func New() *gin.Engine {
	config.LoadAddConfig()

	docs.SwaggerInfo.Version = "1.0"

	r := gin.Default()

	database.Connect()

	r.GET("/api/v1/ping", controller.Ping)
	r.POST("/api/v1/validation_codes", controller.CreateValidationCode)
	r.POST("/api/v1/session", controller.CreateSession)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return r
}
