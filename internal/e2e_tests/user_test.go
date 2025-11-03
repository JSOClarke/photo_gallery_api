package e2e_tests

import (
	"testing"

	"github.com/gin-gonic/gin"
)

func TestCreateUser(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	// userHandler := handlers.UserHandler{}

	r.POST("/api/v1/users/signup")
}
