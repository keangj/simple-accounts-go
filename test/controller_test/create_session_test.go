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

func TestCreateSession(t *testing.T) {
	email := "keangjay@gmail.com"
	code := "1234"
	r := router.New()
	w := httptest.NewRecorder()
	q := database.NewQuery()
	c := context.Background()
	if _, err := q.CreateValidationCode(c, tutorial.CreateValidationCodeParams{
		Email: email,
		Code:  code,
	}); err != nil {
		log.Fatalln(err)
	}
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
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)
	// log.Println(w.Body.String())
	var responseBody struct {
		JWT string `json:"jwt"`
	}
	if err := json.Unmarshal(w.Body.Bytes(), &responseBody); err != nil {
		t.Error(err)
	}
	fmt.Println(w.Body.String())
	assert.Equal(t, 200, w.Code)
}
