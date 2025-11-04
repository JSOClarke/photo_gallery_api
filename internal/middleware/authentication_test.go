package middleware

import (
	"fmt"
	"net/http/httptest"
	"photogallery/internal/models"
	"photogallery/internal/utils"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestAuthenticationMiddleware(t *testing.T) {
	username := "golden_user"
	token, err := utils.CreateToken(username)
	if err != nil {
		t.Error("not able to create token")
	}
	header_value := fmt.Sprint("Bearer ", token)
	// fmt.Println(header_value)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	req := httptest.NewRequest("GET", "/protected", nil)
	req.Header.Set("Authorization", header_value)
	c.Request = req
	Authentication(c)
	val, exists := c.Get("claims")

	assert.True(t, exists)
	assert.Equal(t, username, val.(models.Claims).Username)
}
