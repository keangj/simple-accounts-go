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
	// 从请求中获取 JSON 数据并解析到 requestBody
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	q := database.NewQuery()
	// 查找验证码
	_, err := q.FindValidationCode(c, tutorial.FindValidationCodeParams{
		Email: requestBody.Email,
		Code:  requestBody.Code,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 查找用户, 如果没有则创建用户
	user, err := q.GetUserByEmail(c, requestBody.Email)
	if err != nil {
		user, err = q.CreateUser(c, requestBody.Email)
		if err != nil {
			log.Println("CreateUser fail", err)
			c.String(http.StatusInternalServerError, "Internal Server Error")
			return
		}
	}
	// 生成 JWT
	JWTToken, err := jwt_helper.GenerateJWT(int(user.ID))
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
	c.JSON(http.StatusOK, gin.H{"jwt": JWTToken, "userId": user.ID})
}
