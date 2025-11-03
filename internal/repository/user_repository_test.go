package repository

import (
	"fmt"
	"os"
	"photogallery/internal/models"
	"photogallery/internal/pkg/db"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func InitDB() string {
	err := godotenv.Load("../../.env")
	if err != nil {
		panic("Godotenv failed to load")
	}

	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("PG_HOST"),
		os.Getenv("PG_PORT"),
		os.Getenv("PG_USER"),
		os.Getenv("PG_PASSWORD"),
		os.Getenv("PG_DBNAME"),
	)
	return connStr
}
func TestCreateUser_SUCCESS(t *testing.T) {

	database := db.Connect(InitDB())

	userRep := UserRepo{DB: database}

	username := "user3"
	password_hash := []byte("secureHashedPassword")

	res, err := userRep.CreateUser(username, string(password_hash))

	// teardown before assertion so test fail doesnt stop the teardown

	_, e := database.Exec("DELETE from users WHERE username = $1", res)
	if e != nil {
		t.Fatalf("failed to clean up user: %v", e)

	}
	assert.Equal(t, []byte(username), res)
	assert.NoError(t, err)

}
func TestCreateUser_DUPLICATE(t *testing.T) {

	database := db.Connect(InitDB())

	userRep := UserRepo{DB: database}

	username := "user1"
	password_hash := []byte("secureHashedPassword")
	_, err := database.Exec("INSERT INTO users (username,password_hash) VALUES ($1,$2)", username, password_hash)
	if err != nil {
		t.Fatal("Issue setting up fixture", err)
	}
	_, err1 := userRep.CreateUser(username, string(password_hash))

	// teardown before assertion so test fail doesnt stop the teardown

	_, err = database.Exec("DELETE from users WHERE username = $1", username)
	if err != nil {
		t.Fatalf("failed to clean up user: %v", err)

	}
	fmt.Println("error", err1)
	// assert.Equal(t, []byte(username), res)
	assert.Error(t, err1)

}

// Tests for the happy path with password_hash returned
func TestLoginUser_SUCCESS(t *testing.T) {
	database := db.Connect(InitDB())

	// Fixtures
	fixtures_login := models.LoginRequest{Username: "user1", Password: "securepassword"}

	_, err := database.Exec("INSERT INTO users (username,password_hash) VALUES ($1,$2)", fixtures_login.Username, fixtures_login.Password)
	if err != nil {
		t.Fatal("Fixture opeation failed, testcase no longer valid", err)
	}

	userRepo := UserRepo{DB: database}

	password_hash, err := userRepo.LoginUser(fixtures_login.Username)

	_, err = database.Exec("DELETE from users WHERE username = $1", fixtures_login.Username)
	if err != nil {
		t.Fatalf("failed to clean up user: %v", err)

	}
	fmt.Println("password_hash", string(password_hash))

	assert.NoError(t, err)
}

// Tests for unknown username in the database no rows returned
func TestLoginUser_FAILURE(t *testing.T) {
	database := db.Connect(InitDB())

	// Fixtures
	fixtures_login := models.LoginRequest{Username: "user1"}

	userRepo := UserRepo{DB: database}

	_, err := userRepo.LoginUser(fixtures_login.Username)

	assert.Error(t, err)
	assert.NotEqual(t, nil, err)

	fmt.Println("err", err)
}
