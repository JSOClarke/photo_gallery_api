package router

import (
	"photogallery/internal/handlers"
	"photogallery/internal/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	v1 := r.Group("/api/v1", middleware.Logger())
	v1.GET("/health", handlers.GetHealth)
	v1.POST("/photos",handl)
}
