package services

import (
	"errors"
	"photogallery/internal/repository"
)

type PhotoServiceInterface interface {
	GetPhoto(int) ([]byte, string, error)
	GetAllPhotos([]byte) ([]repository.GetPhotosResponse, error)
}

type PhotoService struct {
	Repo repository.PhotoRepo
}

type NewPhotoService struct {
}

func (ps *PhotoService) GetPhoto(param int) ([]byte, string, error) {

	if param == 20 {
		return []byte("Jordan"), "image/jpeg", nil
	}

	return []byte{}, "", errors.New("No image found")
}

// No business logic really needed here as we passing the username to check the photos against
func (ps *PhotoService) GetAllPhotos(username []byte) ([]repository.GetPhotosResponse, error) {
	data, err := ps.Repo.GetPhotos(username)
	if err != nil {
		return nil, err
	}

	return data, nil
}
