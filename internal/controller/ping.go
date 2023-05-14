package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Ping godoc
// @Summary      ping pong
// @Description  test ping pong
// @Accept       json
// @Produce      json
// @Success      200
// @Failure      400
// @Failure      404
// @Failure      500
// @Router       /ping [get]
func Ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
	log.Println("setupRouter")

}
