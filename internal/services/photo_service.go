package services

import (
	"errors"
	"fmt"
	"log"
	"mime/multipart"
	"os"
	"photogallery/internal/repository"
	"strings"

	"github.com/joho/godotenv"
)

type PhotoServiceInterface interface {
	GetPhoto(int, string) ([]byte, string, error)
	GetAllPhotos([]byte) ([]repository.GetPhotosResponse, error)
	UploadPhotos(files []*multipart.FileHeader) (string, error)
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
	fmt.Println("split_string->", splitString)
	// the e
	if len(splitString) != 2 {
		return nil, "", errors.New("Filename dont provided in correct format")
	}
	contentType := fmt.Sprint(splitString[1], "/", "image")
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

// Takes in the files array and returns the original_filename and ID of the database addition entry
func (ps *PhotoService) UploadPhotos(files []*multipart.FileHeader) (string, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("not able to load env file ", err.Error())
	}

	uploads_path := os.Getenv("UPLOADS_PATH")
	if len(uploads_path) == 0 {
		log.Fatal("uploads path var is 0")
	}

	// var uploaded_files []string

	for _, file := range files {

		split_string := strings.Split(file.Filename, ".")
		if len(split_string) != 2 {
			log.Fatal("cant extract the file type")
		}
		// original_filename := split_string[0]

		if err != nil {
			return "", err
		}
		file_name := fmt.Sprint(uploads_path + "/" + file.Filename)
		out_file, err := os.Create(file_name)
		if err != nil {
			log.Fatal("not able to create a new file iwth os.Create()", err.Error())
		}
		defer out_file.Close()
		var buffer [1024]uint8
		file, err := file.Open()
		if err != nil {
			log.Fatal("not able to open file", err.Error())
		}
		defer file.Close()
		for {
			b_read, err := file.Read(buffer[:])
			if b_read == 0 {
				break
			}
			if err != nil {
				log.Fatal(err.Error())
			}
			out_file.Write(buffer[:b_read])
		}
	}
	return "answer", nil
}
