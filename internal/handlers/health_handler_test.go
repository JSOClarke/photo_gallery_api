package handlers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/gin-gonic/gin"
)

func TestHealth(t *testing.T) {
	r := gin.New()
	r.GET("/api/v1/health", GetHealth)

	req, _ := http.NewRequest("GET", "/api/v1/health", nil)
	// This creates the request ->> wraps the request

	w := httptest.NewRecorder()
	//Records the response for asserts

	// Pass intended request and the recorder to mock the httpWriter
	r.ServeHTTP(w, req)

	fmt.Println("Response", w.Body)
	assert.Equal(t, 200, w.Code)

}
