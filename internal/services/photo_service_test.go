package services

import (
	"photogallery/internal/repository"
	"testing"

	"github.com/stretchr/testify/assert"
)

func (m *MockRepo) GetPhotos(username []byte) ([]repository.GetPhotosResponse, error) {
	m.Called = true
	return []repository.GetPhotosResponse{}, nil
}

func (m *MockRepo) GetPhotoFilename(id int, username string) (*repository.GetPhotoFilenameResponse, error) {
	m.Called = true
	return &repository.GetPhotoFilenameResponse{Hashed_Filename: "test.jpeg"}, nil
}

func TestGetPhoto(t *testing.T) {

	mockRepo := MockRepo{}
	photoService := PhotoService{Repo: &mockRepo}
	image, contentType, err := photoService.GetPhoto(1000, "test")
	if err != nil {
		t.Error("error", err.Error())
	}
	assert.Greater(t, len(image), 100)
	assert.NotEmpty(t, image)

	assert.NotEmpty(t, contentType)
	// fmt.Println("image: ", image)
	// fmt.Println("content type: ", contentType)
}
