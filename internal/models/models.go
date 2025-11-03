package models

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password"  binding:"required"`
}

type PhotosRequest struct {
	Username string `json:"username" binding="required`
}

type Claims struct {
	Username string `json:"username"`
}
