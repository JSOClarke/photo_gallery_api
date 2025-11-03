package middleware

import (
	"net/http"
	"photogallery/internal/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

func Authentication(g *gin.Context) {

	auth_header := g.Request.Header.Get("Authorization")
	if auth_header == "" {
		g.JSON(http.StatusUnauthorized, gin.H{"error": "auth headers not provided"})
		g.Abort()
	}
	token := strings.Split(auth_header, " ")[1]
	mapClaims, err := utils.VerifyJWT(token)
	if err != nil {
		g.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		g.Abort()
	}
	g.Set("claims", mapClaims)

	// add the maps claims to the body and pass to the handler function
	g.Next()
}
