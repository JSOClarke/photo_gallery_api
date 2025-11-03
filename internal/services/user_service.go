package services

import (
	"photogallery/internal/models"
	"photogallery/internal/repository"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type UserServiceInterface interface {
	CreateUser(models.LoginRequest) ([]byte, error)
	LoginUser(models.LoginRequest) ([]byte, error)
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
func (us *UserService) LoginUser(login models.LoginRequest) ([]byte, error) {
	//

	password_hash, err := us.Repo.LoginUser(login.Username)
	if err != nil {
		return nil, err
	}

	err = CompareHashAndPassword(password_hash, []byte(login.Password))
	if err != nil {
		return nil, err
	}
	token, err := CreateToken([]byte(login.Username))
	if err != nil {
		return nil, err
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

func CreateToken(username []byte) ([]byte, error) {
	claims := jwt.MapClaims{"username": username}

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)

	// we need to sign the token with the secret

	jwt_secret := []byte("yourMother")

	signedToken, err := token.SignedString(jwt_secret)
	if err != nil {
		panic("Not able to sign the token")
	}
	return []byte(signedToken), nil
}

func CompareHashAndPassword(password_hash, password []byte) error {
	err := bcrypt.CompareHashAndPassword(password_hash, password)
	if err != nil {
		return err
	}
	return nil
}
