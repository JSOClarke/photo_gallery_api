package repository

import (
	"fmt"
	"log"
	"os"
	"photogallery/internal/pkg/db"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestPhotoRepo(t *testing.T) {
	user := []byte("golden_user")

	if err := godotenv.Load(); err != nil {
		log.Fatal(".Env could not be loaded", err)
	}

	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("PG_HOST"),
		os.Getenv("PG_PORT"),
		os.Getenv("PG_USER"),
		os.Getenv("PG_PASSWORD"),
		os.Getenv("PG_DBNAME"),
	)

	db := db.Connect(connStr)
	pr := PhotoRepo{DB: db}
	responses, err := pr.GetPhotos(user)
	// fmt.Println("response", responses)
	if err != nil {
		t.Error("err", err.Error())
	}
	for _, i := range responses {
		fmt.Println(string(i.Original_file_name))
	}
	assert.Equal(t, 2, len(responses))
}
