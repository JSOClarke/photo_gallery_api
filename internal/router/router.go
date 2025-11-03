package router

import (
	"photogallery/internal/handlers"
	"photogallery/internal/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, userHandler *handlers.UserHandler) {
	v1 := r.Group("/api/v1")
	{
		v1.GET("/health", handlers.GetHealth)
		users := v1.Group("/users")
		{
			users.POST("/signup", userHandler.SignUpUser)
			users.POST("/login", userHandler.LoginUser)
		}
		photos := v1.Group("/photos", middleware.Authentication)
		{
			photos.POST("/upload")

			photos.GET("/images", userHandler.LoginUser)
			photos.GET("/image/:id")
			photos.POST("/transform/:id")

		}
	}

}
