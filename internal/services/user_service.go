package services

import (
	"photogallery/internal/models"
	"photogallery/internal/repository"
	"photogallery/internal/utils"

	"golang.org/x/crypto/bcrypt"
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
	hashed_password, err := hashPassword([]byte(login.Password))
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

	err = CompareHashAndPassword(password_hash, login.Password)
	if err != nil {
		return "", err
	}
	token, err := utils.CreateToken(login.Username)
	if err != nil {
		return "", err
	}
	return token, nil
}

func hashPassword(password []byte) ([]byte, error) {

	hash_cost := 10

	hashed_password, err := bcrypt.GenerateFromPassword(password, hash_cost)
	if err != nil {
		return nil, err
	}
	return hashed_password, nil
}

func CompareHashAndPassword(password_hash, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(password_hash), []byte(password))
	if err != nil {
		return err
	}
	return nil
}
