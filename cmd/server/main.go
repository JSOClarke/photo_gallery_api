package main

import (
	"photogallery/internal/router"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default() // creates the instance of of the engine // This wraps hte middleware, muxer and confiugation settings

	// r.GET("/health", func(ctx *gin.Context) { // this acts like a  blanket for the specific htptp method and the path that is requested with it
	// 	ctx.JSON(http.StatusOK, gin.H{"message": "Your boys eat butt"})
	// })

	router.RegisterRoutes(r)
	r.Run()
}
