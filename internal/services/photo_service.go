package services

import (
	"errors"
	"fmt"
	"os"
	"photogallery/internal/repository"
	"strings"

	"github.com/joho/godotenv"
)

type PhotoServiceInterface interface {
	GetPhoto(int, string) ([]byte, string, error)
	GetAllPhotos([]byte) ([]repository.GetPhotosResponse, error)
}

type PhotoService struct {
	Repo repository.PhotoRepoInterface
}

func NewPhotoService(repo *repository.PhotoRepo) *PhotoService {
	return &PhotoService{Repo: repo}
}

func (ps *PhotoService) GetPhoto(id int, username string) ([]byte, string, error) {

	repo_result, err := ps.Repo.GetPhotoFilename(id, username)
	if err != nil {
		return nil, "", err
	}

	err = godotenv.Load()
	if err != nil {
		return nil, "", err

	}

	base_path := os.Getenv("UPLOADS_PATH")

	// we need to get the file from the upload folder
	path := fmt.Sprint(base_path + "/" + repo_result.Hashed_Filename)
	splitString := strings.Split(repo_result.Hashed_Filename, ".")
	// the e
	if len(splitString) != 2 {
		return nil, "", errors.New("Filename dont provided in correct format")
	}
	contentType := splitString[1]
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, "", err
	}
	return content, contentType, nil
}

// No business logic really needed here as we passing the username to check the photos against
func (ps *PhotoService) GetAllPhotos(username []byte) ([]repository.GetPhotosResponse, error) {
	data, err := ps.Repo.GetPhotos(username)
	if err != nil {
		return nil, err
	}

	return data, nil
}
