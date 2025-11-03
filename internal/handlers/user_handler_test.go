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

	token := []byte(`"token":"tokensystem"`)

	if login.Username == "token" {
		return token, us.Err
	}
	return []byte(login.Username), us.Err
}

func (us *MockService) LoginUser(login models.LoginRequest) ([]byte, error) {
	us.Called = true
	if login.Username == "fail" {
		return nil, errors.New("Name not found in the database")
	}

	token := []byte(`"token":"tokensystem"`)

	if login.Username == "token" {
		return token, nil
	}
	return []byte(login.Username), nil
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

func TestLoginUser_SUCCESS(t *testing.T) {
	gin.SetMode(gin.TestMode)

	loginUser := "/api/v1/users/login"

	r := gin.Default()

	mockService := MockService{Err: nil}

	userHandler := UserHandler{&mockService}
	r.POST(loginUser, userHandler.LoginUser)

	w := httptest.NewRecorder()
	login := models.LoginRequest{Username: "token", Password: "securepassword"}

	json_body, err := json.Marshal(login)
	if err != nil {
		t.Fatal("Error marshalling the login", err)
	}

	req, err := http.NewRequest(http.MethodPost, loginUser, bytes.NewReader(json_body))

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusAccepted, w.Code)
	assert.Contains(t, w.Body.String(), "token")

}

func TestLoginUser_FAILURE(t *testing.T) {
	gin.SetMode(gin.TestMode)

	loginUser := "/api/v1/users/login"

	r := gin.Default()

	mockService := MockService{Err: nil}

	userHandler := UserHandler{&mockService}
	r.POST(loginUser, userHandler.LoginUser)

	w := httptest.NewRecorder()
	login := models.LoginRequest{Username: "fail", Password: "securepassword"}

	json_body, err := json.Marshal(login)
	if err != nil {
		t.Fatal("Error marshalling the login", err)
	}

	req, err := http.NewRequest(http.MethodPost, loginUser, bytes.NewReader(json_body))
	if err != nil {
		t.Fatal("Not able to form the Request")
	}

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
	assert.Contains(t, w.Body.String(), "error")
	// fmt.Println("error", w.Body.String())

	// assert.Contains(t, w.Body.String(), "token")

}

func TestLoginUser_EMPTY_BODY(t *testing.T) {
	gin.SetMode(gin.TestMode)

	loginUser := "/api/v1/users/login"

	r := gin.Default()

	mockService := MockService{Err: nil}

	userHandler := UserHandler{&mockService}
	r.POST(loginUser, userHandler.LoginUser)

	w := httptest.NewRecorder()

	req, err := http.NewRequest(http.MethodPost, loginUser, nil)
	if err != nil {
		t.Fatal("Not able to form the Request")
	}

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	// Coming back with request body is empty but different formatting fix this tommorrow
	assert.Contains(t, "Request body is empty", w.Body.String())
	// assert.Contains(t, w.Body.String(), "error")
	fmt.Println("error", w.Body.String())

	// assert.Contains(t, w.Body.String(), "token")

}
