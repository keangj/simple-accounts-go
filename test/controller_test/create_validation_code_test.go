package controller_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"simple-accounts/internal/database"
	"simple-accounts/internal/router"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateValidationCode(t *testing.T) {
	email := "keangjay@gmail.com"
	// Setup
	r := router.New()
	// Assertions
	w := httptest.NewRecorder()

	q := database.NewQuery()
	count1, _ := q.CountValidationCode(context.Background(), email)
	req, _ := http.NewRequest(
		"POST",
		"/api/v1/validation_codes",
		strings.NewReader(`{"email":"`+email+`"}`),
	)
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)
	count2, _ := q.CountValidationCode(context.Background(), email)
	assert.Equal(t, count2-1, count1)
	assert.Equal(t, 200, w.Code)
}
