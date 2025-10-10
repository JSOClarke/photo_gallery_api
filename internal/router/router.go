package router

import (
	"photogallery/internal/handlers"
	"photogallery/internal/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, userHandler *handlers.UserHandler) {
	v1 := r.Group("/api/v1", middleware.Logger())
	{
		v1.GET("/health", handlers.GetHealth)
		users := v1.Group("/users")
		{
			users.POST("", userHandler.SignUpUser)
		}
	}

}
