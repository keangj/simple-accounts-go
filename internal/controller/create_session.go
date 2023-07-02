package controller

import (
	"net/http"
	"simple-accounts/config/tutorial"
	"simple-accounts/internal/database"

	"github.com/gin-gonic/gin"
)

func CreateSession(c *gin.Context) {
	var requestBody struct {
		Email string `json:"email" binding:"required"`
		Code  string `json:"code" binding:"required"`
	}
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	q := database.NewQuery()
	_, err := q.FindValidationCode(c, tutorial.FindValidationCodeParams{
		Email: requestBody.Email,
		Code:  requestBody.Code,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	JWTToken := "jwt"
	// responseBody := struct {
	// 	JWT string `json:"jwt"`
	// }{
	// 	JWT: JWTToken,
	// }
	// c.JSON(http.StatusOK, responseBody)
	c.JSON(http.StatusOK, gin.H{"JWT": JWTToken})
}
