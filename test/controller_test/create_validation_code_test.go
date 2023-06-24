package controller_test

import (
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"simple-accounts/config"
	"simple-accounts/internal/router"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateValidationCode(t *testing.T) {
	config.LoadAddConfig()
	log.Println("test ----------")
	pwd, _ := os.Getwd()
	log.Println(pwd)
	// Setup
	r := router.New()
	// Assertions
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(
		"POST",
		"/api/v1/validation_codes",
		strings.NewReader(`{"email":"keangj@outlook.com"}`),
	)
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)
	// fmt.Println(w.Body.String(), w.Code)
	// log.Fatalln(w.Body.String(), w.Code)
	assert.Equal(t, 200, w.Code)
	// assert.Equal(t, "pong", w.Body.String())
}
