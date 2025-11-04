package handlers

import (
	"errors"
	"fmt"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"photogallery/internal/models"
	"photogallery/internal/repository"
	"photogallery/internal/services"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
)

type MockPhotoService struct {
	Called bool
}

const SUCCESS = 7
const FAILURE = 666

func (ps *MockPhotoService) GetPhoto(param int, username string) ([]byte, string, error) {
	if param == SUCCESS {
		return []byte("Jordan"), "image/jpeg", nil
	}

	return []byte{}, "", errors.New("No image found")
}
func (ps *MockPhotoService) UploadPhotos(files []*multipart.FileHeader) (string, error) {
	return "done", nil
}

func (ps *MockPhotoService) GetAllPhotos(username []byte) ([]repository.GetPhotosResponse, error) {
	ps.Called = true
	if string(username) == "fail" {
		return []repository.GetPhotosResponse{}, errors.New("Service error")
	}
	entry := repository.GetPhotosResponse{Original_file_name: "long_drive"}
	return []repository.GetPhotosResponse{entry}, nil
}

func TestGetPhoto(t *testing.T) {
	// r := gin.Default()
	mockService := MockPhotoService{}
	photoHandlerx := PhotoHandler{Service: &mockService}

	// r.GET("/testGetPhoto/:id", photoHandlerx.GetPhoto)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{gin.Param{Key: "id", Value: fmt.Sprint(SUCCESS)}}
	claims := models.Claims{
		Username: "Jordan",
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}
	c.Set("claims", claims)
	req := httptest.NewRequest("GET", "/testGetPhoto/1", nil)
	c.Request = req
	//	The body response with the image
	// The fake write object that will record what is being written inside like the body etc
	// r.ServeHTTP(w, req)
	photoHandlerx.GetPhoto(c)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "image/jpeg", w.Header().Get("Content-type"))
	fmt.Println("image", w.Body)

}
func TestGetPhoto_Failure(t *testing.T) {
	// r := gin.Default()
	mockService := MockPhotoService{}
	photoHandlerx := PhotoHandler{Service: &mockService}

	// r.GET("/testGetPhoto/:id", photoHandlerx.GetPhoto)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{gin.Param{Key: "id", Value: fmt.Sprint(FAILURE)}}
	req := &http.Request{}
	c.Request = req
	//	The body response with the image
	// The fake write object that will record what is being written inside like the body etc
	// r.ServeHTTP(w, req)
	photoHandlerx.GetPhoto(c)

	assert.Equal(t, 400, w.Code)
	// assert.Equal(t, "image/jpeg", w.Header().Get("Content-type"))

}

func TestGetAllPhoto_NoClaims_Failure(t *testing.T) {
	r := gin.Default()
	mockService := MockPhotoService{}

	photoHandlerx := PhotoHandler{Service: &mockService}
	endpoint := "/testGetAllPhotos"
	r.GET(endpoint, photoHandlerx.GetAllPhotos)
	w := httptest.NewRecorder()

	req, err := http.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		panic("Not able to create request, test invalid")
	}
	//	The body response with the image
	// The fake write object that will record what is being written inside like the body etc
	r.ServeHTTP(w, req)
	// fmt.Println("error", w.Body.String())
	assert.Equal(t, 400, w.Code)
	assert.Contains(t, w.Body.String(), "claims do not exist")
}
func TestGetAllPhoto_SUCCESS(t *testing.T) {
	mockService := MockPhotoService{}
	photoHandlerx := PhotoHandler{Service: &mockService}

	claims := models.Claims{
		Username: "Jordan",
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Set("claims", claims)
	c.Request = &http.Request{}
	photoHandlerx.GetAllPhotos(c)

	assert.Equal(t, http.StatusOK, w.Code)
	fmt.Println("Response:", w.Body.String())
}

func TestPhotoHandler_UploadPhoto(t *testing.T) {

	tests := []struct {
		name string // description of this test case
		// Named input parameters for receiver constructor.
		service *services.PhotoService
		// Named input parameters for target function.
		c *gin.Context
	}{
		// TODO: Add test cases.

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ph := NewPhotoHandler(tt.service)
			ph.UploadPhoto(tt.c)
		})
	}
}
