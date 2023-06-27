package controller

import (
	"crypto/rand"
	"log"
	"net/http"
	"simple-accounts/config/tutorial"
	"simple-accounts/internal/database"
	"simple-accounts/internal/email"

	"github.com/gin-gonic/gin"
)

type Body struct {
	Email string `json:"email" binding:"required,email"`
}

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
	var body Body
	if err := c.ShouldBindJSON(&body); err != nil {
		c.String(400, "error")
		return
	}
	str, err := generateDigits()
	if err != nil {
		log.Println("[generateDigits fail]", err)
		c.String(500, "生成验证码失败")
		return
	}
	q := database.NewQuery()
	vc, err := q.CreateValidationCode(c, tutorial.CreateValidationCodeParams{
		Email: body.Email,
		Code:  str,
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

func generateDigits() (string, error) {
	len := 4
	b := make([]byte, len)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	digits := make([]byte, len)
	for i := range b {
		digits[i] = b[i]%10 + 48
	}
	return string(digits), nil
}
