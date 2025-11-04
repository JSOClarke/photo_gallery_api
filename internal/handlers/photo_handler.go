package handlers

import (
	"net/http"
	"photogallery/internal/models"
	"photogallery/internal/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PhotoHandler struct {
	Service services.PhotoServiceInterface
}

func NewPhotoHandler(service *services.PhotoService) *PhotoHandler {
	return &PhotoHandler{Service: service}
}

func (ph *PhotoHandler) GetAllPhotos(c *gin.Context) {

	claims, exist := c.Get("claims")

	if !exist {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Token is missing or expired"})
	}
	username := claims.(models.Claims).Username
	if username == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username missing from jwt claim"})
		return
	}
	service, err := ph.Service.GetAllPhotos([]byte(username))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, service)
}

// Gets the given photo from the database by provided ID pararm
func (ph *PhotoHandler) GetPhoto(c *gin.Context) {
	id, isFound := c.Params.Get("id")

	if !isFound {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Incorrect parameter passed"})
		return
	}
	id_int, err := strconv.Atoi(id)

	claims, exist := c.Get("claims")

	if !exist {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Token is missing or expired"})
	}
	username := claims.(models.Claims).Username
	if username == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username missing from jwt claim"})
		return
	}

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	data, contentType, err := ph.Service.GetPhoto(id_int, username)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//Set Headers for the binary data sent back
	c.Writer.Header().Set("Content-type", contentType)
	c.Data(http.StatusOK, contentType, data)

}
