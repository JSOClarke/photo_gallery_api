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

type PhotoRepoInterface interface {
	GetPhotos(username []byte) ([]GetPhotosResponse, error)
}

type GetPhotosResponse struct {
	Original_file_name []byte
}

func (pr *PhotoRepo) GetPhotos(username []byte) ([]GetPhotosResponse, error) {
	rows, err := pr.DB.Query("SELECT original_filename FROM photos WHERE owned_by=$1", username)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var response_array []GetPhotosResponse

	for rows.Next() {
		var query_response GetPhotosResponse

		if err := rows.Scan(&query_response.Original_file_name); err != nil {
			return nil, err
		}
		response_array = append(response_array, query_response)
	}
	return response_array, nil
}
