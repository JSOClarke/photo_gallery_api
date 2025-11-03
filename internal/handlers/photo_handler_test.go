package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"photogallery/internal/repository"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type MockPhotoService struct {
	Called bool
}

func (ps *MockPhotoService) GetPhoto(param int) ([]byte, string, error) {
	if param == 20 {
		return []byte("Jordan"), "image/jpeg", nil
	}

	return []byte{}, "", errors.New("No image found")
}

func (ps *MockPhotoService) GetAllPhotos(username []byte) ([]repository.GetPhotosResponse, error) {
	ps.Called = true
	if string(username) == "fail" {
		return []repository.GetPhotosResponse{}, errors.New("Service error")
	}
	entry := repository.GetPhotosResponse{Original_file_name: []byte("long_drive")}
	return []repository.GetPhotosResponse{entry}, nil
}

func TestGetPhoto_SUCCESS(t *testing.T) {
	r := gin.Default()
	mockService := MockPhotoService{}

	photoHandlerx := PhotoHandler{Service: &mockService}

	r.GET("/testGetPhoto/:id", photoHandlerx.GetPhoto)
	w := httptest.NewRecorder()

	req, err := http.NewRequest(http.MethodGet, "/testGetPhoto/20", nil)
	if err != nil {
		panic("Not able to create request, test invalid")
	}
	//	The body response with the image
	// The fake write object that will record what is being written inside like the body etc
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "image/jpeg", w.Header().Get("Content-type"))

}
func TestGetPhoto_FAILURE(t *testing.T) {
	r := gin.Default()
	mockService := MockPhotoService{}

	photoHandlerx := PhotoHandler{Service: &mockService}

	r.GET("/testGetPhoto/:id", photoHandlerx.GetPhoto)
	w := httptest.NewRecorder()

	req, err := http.NewRequest(http.MethodGet, "/testGetPhoto/101", nil)
	if err != nil {
		panic("Not able to create request, test invalid")
	}
	//	The body response with the image
	// The fake write object that will record what is being written inside like the body etc
	r.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)

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
	r := gin.Default()
	mockService := MockPhotoService{}

	photoHandlerx := PhotoHandler{Service: &mockService}
	endpoint := "/testGetAllPhotos"
	r.GET(endpoint, photoHandlerx.GetAllPhotos)
	w := httptest.NewRecorder()

	c, _ := gin.CreateTestContext(w)

	req, err := http.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		panic("Not able to create request, test invalid")
	}
	c.Request = req

	// c.Set()

	//	The body response with the image
	// The fake write object that will record what is being written inside like the body etc
	r.ServeHTTP(w, req)
	// fmt.Println("error", w.Body.String())
	assert.Equal(t, 200, w.Code)
	fmt.Println(w.Body)
	// assert.Contains(t, w.Body.String(), "claims do not exist")
}
