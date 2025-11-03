package middleware

import (
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestAuthenticationMiddleware(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	req := httptest.NewRequest("GET", "/protected", nil)
	req.Header.Set("Authorization", "Bearer ZXlKaGJHY2lPaUpJVXpVeE1pSXNJblI1Y0NJNklrcFhWQ0o5LmV5SjFjMlZ5Ym1GdFpTSTZJbG95T1hOYVIxWjFXRE5XZWxwWVNUMGlmUS4yTnk1TEhjaXFGWlVJLVV6VVRBakRxbUxXZllJMFB1NFp2VXJYZGl0Z3NzdUV0U0dCSU1vVHZxZ3lBOWp1TGxDdmVlODJzNElERF9LOGdyaFQzb2dCZw==")
	c.Request = req
	Authentication(c)
	val, exists := c.Get("claims")
	assert.True(t, exists)
	fmt.Println("val", val)
}
