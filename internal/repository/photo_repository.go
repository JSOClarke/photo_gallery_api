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
	GetPhotoFilename(id int, username string) (*GetPhotoFilenameResponse, error)
}

type GetPhotosResponse struct {
	Original_file_name string
}

type GetPhotoFilenameResponse struct {
	Hashed_Filename string
}

func (pr *PhotoRepo) GetPhotoFilename(id int, username string) (*GetPhotoFilenameResponse, error) {
	row := pr.DB.QueryRow("SELECT Hashed_Filename FROM photos WHERE owned_by=$1 AND id=$2", username, id)

	var response GetPhotoFilenameResponse
	err := row.Scan(&response.Hashed_Filename)
	if err != nil {
		if err == sql.ErrNoRows {
			return &GetPhotoFilenameResponse{}, nil
		}
		return nil, err
	}
	return &response, nil
}

func (pr *PhotoRepo) GetPhotos(username []byte) ([]GetPhotosResponse, error) {
	rows, err := pr.DB.Query("SELECT Hashed_Filename FROM photos WHERE owned_by=$1", username)
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

	if response_array == nil {
		response_array = []GetPhotosResponse{}
	}
	return response_array, nil

}
