package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"photogallery/internal/models"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// TO implement datadriven

type MockService struct {
	Called bool
	Err    error
}

func (us *MockService) CreateUser(login models.LoginRequest) ([]byte, error) {
	us.Called = true
	if login.Username == "fail" {
		return nil, errors.New("Unique name issue, duplicate entry")
	}
	return []byte(login.Username), us.Err
}

func TestSignUpUser_SUCCESS(t *testing.T) {
	gin.SetMode(gin.TestMode)
	signupEndpoint := "/api/v1/users"
	r := gin.New()

	// inject the dependecny for testing

	mockService := &MockService{Err: nil}

	userHandler := UserHandler{Service: mockService}

	// user the injected handler as the endpoints.
	r.POST("/api/v1/users", userHandler.SignUpUser)

	w := httptest.NewRecorder()

	body := models.LoginRequest{Username: "Jordan", Password: "Place"}
	body_bytes, _ := json.Marshal(body)
	req, _ := http.NewRequest("POST", signupEndpoint, bytes.NewReader(body_bytes))

	req.Header.Set("Content-Type", "application/json")
	fmt.Println("res", req.Body)

	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Contains(t, w.Body.String(), body.Username)

}

func TestSignUpUser_FAILURE(t *testing.T) {
	gin.SetMode(gin.TestMode)
	signupEndpoint := "/api/v1/users"
	r := gin.New()

	// inject the dependecny for testing

	mockService := &MockService{Err: nil}

	userHandler := UserHandler{Service: mockService}

	// user the injected handler as the endpoints.
	r.POST("/api/v1/users", userHandler.SignUpUser)

	w := httptest.NewRecorder()

	body := models.LoginRequest{Username: "fail", Password: "Place"}
	body_bytes, _ := json.Marshal(body)
	req, _ := http.NewRequest("POST", signupEndpoint, bytes.NewReader(body_bytes))

	req.Header.Set("Content-Type", "application/json")

	r.ServeHTTP(w, req)

	assert.Equal(t, 500, w.Code)
}

func TestLoginUser()
