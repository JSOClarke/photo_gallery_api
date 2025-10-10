package handlers

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password"  binding:"required"`
}

type UserHandler struct {
	DB *sql.DB
}

// Prevalidation using validator

func NewUserHandler(db *sql.DB) *UserHandler {
	return &UserHandler{DB: db}
}

func (uh *UserHandler) SignUpUser(r *gin.Context) {
	// Grab the username and passsword hash form here
	// validation on the usernma

	var user LoginRequest

	if err := r.ShouldBindJSON(&user); err != nil {
		r.JSON(http.StatusBadRequest, gin.H{"Error": "Body is missing paramters"})

	}

	// if user.Username == "" || user.Password == "" {
	// 	r.JSON(http.StatusBadRequest, gin.H{"error": "Missing parameters"})
	// }

	fmt.Println("Recieved user: ", user)
	r.JSON(http.StatusOK, gin.H{"message": "Succesfully created the user not on database yet though"})
}
