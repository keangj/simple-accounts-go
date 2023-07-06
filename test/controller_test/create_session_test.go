package controller_test

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"simple-accounts/config/tutorial"
	"simple-accounts/internal/database"
	"simple-accounts/internal/router"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var (
	r *gin.Engine
	q *tutorial.Queries
	c context.Context
)

func setupTest(t *testing.T) func(t *testing.T) {
	r = router.New()
	q = database.NewQuery()
	c = context.Background()
	if err := q.DeleteAllUsers(c); err != nil {
		t.Fatal(err)
	}
	return func(t *testing.T) {
		database.Close()
	}
}

func TestCreateSession(t *testing.T) {
	teardownTest := setupTest(t)
	defer teardownTest(t)

	// 创建验证码
	email := "123@gmail.com"
	code := "1234"
	_, err := q.CreateValidationCode(c, tutorial.CreateValidationCodeParams{
		Email: email,
		Code:  code,
	})
	if err != nil {
		log.Fatalln(err)
	}
	// 创建用户
	user, err := q.CreateUser(c, email)
	if err != nil {
		log.Fatalln(err)
	}
	// 创建请求
	w := httptest.NewRecorder()
	j := gin.H{
		"email": email,
		"code":  code,
	}
	bytes, _ := json.Marshal(j)
	req, _ := http.NewRequest(
		"POST",
		"/api/v1/session",
		strings.NewReader(string(bytes)),
	)
	req.Header.Set("Content-Type", "application/json") // 设置请求头
	r.ServeHTTP(w, req)                                // 发送请求
	// log.Println(w.Body.String())
	// 解析响应
	var responseBody struct {
		JWT    string `json:"jwt"`
		UserId int32  `json:"userId"`
	}
	if err := json.Unmarshal(w.Body.Bytes(), &responseBody); err != nil {
		t.Error(err)
	}
	fmt.Println(w.Body.String())
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, user.ID, responseBody.UserId)
}
