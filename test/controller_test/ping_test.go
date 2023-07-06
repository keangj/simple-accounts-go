package controller_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"simple-accounts/internal/router"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPing(t *testing.T) {
	// Setup
	r := router.New()

	// Assertions
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	r.ServeHTTP(w, req)
	fmt.Println(w.Body.String(), w.Code)
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "pong", w.Body.String())
}
