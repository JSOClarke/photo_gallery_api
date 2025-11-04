package services

import (
	"photogallery/internal/models"
	"photogallery/internal/repository"
	"photogallery/internal/utils"
)

type UserServiceInterface interface {
	CreateUser(models.LoginRequest) ([]byte, error)
	LoginUser(models.LoginRequest) (string, error)
}

type UserService struct {
	Repo repository.UserRepoInterface
}

func NewUserService(repo *repository.UserRepo) *UserService {
	return &UserService{Repo: repo}
}

func (us *UserService) CreateUser(login models.LoginRequest) ([]byte, error) {
	hashed_password, err := utils.HashBinaryData([]byte(login.Password), 10)
	if err != nil {
		return nil, err
	}

	// call the repo layer
	repo_response, err := us.Repo.CreateUser(login.Username, string(hashed_password))
	if err != nil {
		return nil, err
	}
	return repo_response, nil
}

// Returns token when check against username and password in DB
func (us *UserService) LoginUser(login models.LoginRequest) (string, error) {
	//

	password_hash, err := us.Repo.LoginUser(login.Username)
	if err != nil {
		return "", err
	}

	err = utils.CompareHashAndPassword(password_hash, login.Password)
	if err != nil {
		return "", err
	}
	token, err := utils.CreateToken(login.Username)
	if err != nil {
		return "", err
	}
	return token, nil
}
