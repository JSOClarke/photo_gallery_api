package repository

import (
	"database/sql"
)

type UserRepo struct {
	DB *sql.DB
}

type UserRepoInterface interface {
	CreateUser(username, password string) ([]byte, error)
	LoginUser(username string) (string, error)
}

func NewRepoService(s *sql.DB) *UserRepo {
	return &UserRepo{DB: s}
}

func (us *UserRepo) CreateUser(username, password_hash string) ([]byte, error) {

	row := us.DB.QueryRow("INSERT INTO users (username,password_hash) VALUES ($1,$2) RETURNING username", username, password_hash)

	// There will probably be a specific reason for this that we should be sending back to the handlers and the client. eg if there is a unqique enitity issue (duplicates)

	var res []byte

	err := row.Scan(&res)
	if err != nil {
		return nil, err
	}

	return res, nil
}
func (us *UserRepo) LoginUser(username string) (string, error) {

	row := us.DB.QueryRow("SELECT password_hash from users where username=$1", username)

	var password_hash string

	err := row.Scan(&password_hash)
	if err != nil {
		return "", err
	}

	return password_hash, nil
}
