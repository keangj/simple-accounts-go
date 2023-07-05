package controller

import (
	"log"
	"net/http"
	"simple-accounts/config/tutorial"
	"simple-accounts/internal/database"
	"simple-accounts/internal/jwt_helper"

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
	JWTToken, err := jwt_helper.GenerateJWT(1)
	if err != nil {
		log.Println(err)
		c.String(http.StatusInternalServerError, "Internal Server Error")
		return
	}
	// responseBody 简化为 gin.H{"JWT": JWTToken}
	// responseBody := struct {
	// 	JWT string `json:"jwt"`
	// }{
	// 	JWT: JWTToken,
	// }
	// c.JSON(http.StatusOK, responseBody)
	c.JSON(http.StatusOK, gin.H{"JWT": JWTToken})
}
