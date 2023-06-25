package controller

import (
	"log"
	"net/http"
	"simple-accounts/config/tutorial"
	"simple-accounts/internal/database"
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
	q := database.NewQuery()
	vc, err := q.CreateValidationCode(c, tutorial.CreateValidationCodeParams{
		Email: body.Email,
		Code:  "666666",
	})
	if err != nil {
		c.Status(400)
		return
	}
	if err := email.SendValidateCode(vc.Email, vc.Code); err != nil {
		log.Println(err)
		c.String(500, "send error")
		return
	}
	log.Println(body.Email)
	c.Status(http.StatusOK)
}
