package repository

import (
	"fmt"
	"os"
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
