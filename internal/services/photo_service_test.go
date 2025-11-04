package services

import (
	"fmt"
	"photogallery/internal/repository"
	"testing"
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
	image, contentType, err := photoService.GetPhoto(2, "golden_user")
	if err != nil {
		t.Error("error", err.Error())
	}
	fmt.Println("image: ", image)
	fmt.Println("content type: ", contentType)
}
