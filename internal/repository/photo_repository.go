package repository

import (
	"database/sql"
)

type PhotoRepo struct {
	DB *sql.DB
}

func NewPhotoRepo(db *sql.DB) *PhotoRepo {
	return &PhotoRepo{DB: db}
}

type GetPhotosResponse struct {
	Original_file_name []byte
}

func (pr *PhotoRepo) GetPhotos(username []byte) ([]GetPhotosResponse, error) {
	rows, err := pr.DB.Query("SELECT original_file_name FROM photos WHERE owned_by=$1", username)
	if err != nil {
		return nil, err
	}
	// the array
	defer rows.Close()
	var response_array []GetPhotosResponse
	var query_response GetPhotosResponse

	for rows.Next() {
		rows.Scan(&query_response)
		response_array = append(response_array, query_response)

	}
	return response_array, nil
}
