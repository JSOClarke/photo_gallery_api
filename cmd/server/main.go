package main

import (
	"fmt"
	"log"
	"os"
	"photogallery/internal/handlers"
	"photogallery/internal/pkg/db"
	"photogallery/internal/repository"
	"photogallery/internal/router"
	"photogallery/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func ServerSetup(st *handlers.UserHandler) *gin.Engine {
	r := gin.Default()
	router.RegisterRoutes(r, st)
	return r
}

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal(".Env could not be loaded", err)
	}

	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("PG_HOST"),
		os.Getenv("PG_PORT"),
		os.Getenv("PG_USER"),
		os.Getenv("PG_PASSWORD"),
		os.Getenv("PG_DBNAME"),
	)

	database := db.Connect(connStr)
	defer database.Close()
	userRep := repository.NewRepoService(database)
	userService := services.NewUserService(userRep)
	handlers := handlers.NewUserHandler(userService)

	r := ServerSetup(handlers)
	r.Run(":4001")
}
