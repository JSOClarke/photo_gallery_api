package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetHealth(r *gin.Context) {
	r.JSON(http.StatusOK, gin.H{"Message": "Please dont talk to me"})
	fmt.Println("Inside hte handler")

	// How does it
}
