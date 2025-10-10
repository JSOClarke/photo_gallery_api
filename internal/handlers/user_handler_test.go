package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"photogallery/internal/handlers"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestSignUpUser_SUCCESS(t *testing.T) {

	signupEndpoint := "/api/v1/users"
	r := gin.New()

	r.POST("/api/v1/users", SignUpUser)

	w := httptest.NewRecorder()

	body := LoginRequest{Username: "Alice", Password: "different"}
	body_bytes, _ := json.Marshal(body)
	req, _ := http.NewRequest("POST", signupEndpoint, bytes.NewReader(body_bytes))

	req.Header.Set("Content-Type", "application/json")

	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestSignUpUser_FAILURE(t *testing.T) {

	signupEndpoint := "/api/v1/users"
	r := gin.New()

	userHandler := handlers.UserHandler{}
	r.POST("/api/v1/users")

	w := httptest.NewRecorder()

	body := LoginRequest{Username: "Alice"}

	body_bytes, _ := json.Marshal(body)
	req, _ := http.NewRequest("POST", signupEndpoint, bytes.NewReader(body_bytes))

	req.Header.Set("Content-Type", "application/json")

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}
