package handlers

import (
	"net/http"
	"photogallery/internal/models"
	"photogallery/internal/services"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	Service services.UserServiceInterface
}

// Prevalidation using validator

func NewUserHandler(service *services.UserService) *UserHandler {
	return &UserHandler{Service: service}
}

func (uh *UserHandler) SignUpUser(r *gin.Context) {

	var user models.LoginRequest

	if err := r.ShouldBindJSON(&user); err != nil {
		r.JSON(http.StatusBadRequest, gin.H{"Error": "Body is missing paramters"})

	}

	service, err := uh.Service.CreateUser(user)

	if err != nil {
		r.JSON(http.StatusInternalServerError, gin.H{"Error": err})
		return
	}
	r.JSON(http.StatusOK, gin.H{"username": string(service)})
}
