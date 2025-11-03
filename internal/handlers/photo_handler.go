package handlers

import (
	"encoding/json"
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

	val, exists := c.Get("claims")
	if exists == false {
		c.JSON(http.StatusBadRequest, gin.H{"error": "token claims do not exist"})
		return
	}
	claims := val.(models.Claims)
	username := claims.Username
	service, err := ph.Service.GetAllPhotos([]byte(username))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	service_json, err := json.Marshal(service)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, service_json)
}

// Gets the given photo from the database by provided ID pararm
func (ph *PhotoHandler) GetPhoto(c *gin.Context) {
	id, isFound := c.Params.Get("id")

	if isFound == false {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Incorrect parameter passed"})
		return
	}

	id_int, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	data, contentType, err := ph.Service.GetPhoto(id_int)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//Set Headers for the binary data sent back
	c.Writer.Header().Set("Content-type", contentType)
	c.Data(http.StatusOK, contentType, data)

}

func (ph *PhotoHandler) GetPhotos(c *gin.Context) {

}
