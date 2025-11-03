package services

import (
	"errors"
	"photogallery/internal/models"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

type MockRepo struct {
	Called bool
	Error  error
}

func (mr *MockRepo) CreateUser(username, password string) ([]byte, error) {
	mr.Called = true
	if username == "fail" {
		return nil, errors.New("unique name issue,duplicate entry")
	}

	return []byte(username), mr.Error
}

func (mr *MockRepo) LoginUser(username string) ([]byte, error) {
	mr.Called = true
	if username == "fail" {
		return nil, errors.New("username not found")
	}
	token, err := hashPassword([]byte("securepassword"))
	if err != nil {
		return nil, err
	}
	return token, nil
}

func TestCreateUser_SUCCESS(t *testing.T) {

	// We are going tobe testing the validation and the hashing of the password basically any log ic before the db layer
	mockRepo := MockRepo{}
	service := UserService{Repo: &mockRepo}

	loginReq := models.LoginRequest{Username: "Leonard", Password: "Hoffstater"}

	res, err := service.CreateUser(loginReq)

	assert.Equal(t, res, []byte(loginReq.Username))
	assert.True(t, mockRepo.Called, true)
	assert.Equal(t, err, nil)
	// In will be a valid username and password and out will be the repo message depending.

}

func TestCreateUser_FAILURE(t *testing.T) {
	// We are going tobe testing the validation and the hashing of the password basically any log ic before the db layer
	mockRepo := MockRepo{}
	service := UserService{Repo: &mockRepo}

	loginReq := models.LoginRequest{Username: "fail", Password: "Hoffstater"}

	res, err := service.CreateUser(loginReq)

	assert.Equal(t, res, []byte(nil))
	assert.True(t, mockRepo.Called, true)
	assert.Error(t, err)
	assert.NotEqual(t, nil, err)

	// In will be a valid username and password and out will be the repo message depending.

}

func TestLoginUser_SUCCESS(t *testing.T) {

	mockRepo := MockRepo{}
	userService := UserService{Repo: &mockRepo}

	login := models.LoginRequest{Username: "user1", Password: "securepassword"}

	token, err := userService.LoginUser(login)
	assert.NoError(t, err)
	assert.GreaterOrEqual(t, len(token), 1)
}

func TestPasswordHash(t *testing.T) {

	password_plain := "securepassword"
	res, _ := hashPassword([]byte(password_plain))

	assert.NotEqual(t, password_plain, res)
	assert.NoError(t, bcrypt.CompareHashAndPassword(res, []byte(password_plain)))
}

func TestCreateToken(t *testing.T) {

	username := []byte("Jordan")

	token, err := CreateToken(username)

	assert.NoError(t, err)
	// fmt.Println("token", string(token))
	assert.NotEqual(t, token, username)
}
