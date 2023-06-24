package controller

import (
	"log"
	"net/http"
	"simple-accounts/internal/email"

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
func CreateValidationCode(c *gin.Context) {
	var body struct {
		Email string `json:"email" binding:"required,email"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.String(400, "error")
		return
	}
	if err := email.SendValidateCode(body.Email, "666666"); err != nil {
		log.Println(err)
		c.String(500, "send error")
		return
	}
	log.Println(body.Email)
	c.Status(http.StatusOK)
}
