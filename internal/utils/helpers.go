package utils

import (
	"errors"
	"fmt"
	"os"
	"photogallery/internal/models"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

// Create a model and tests for this.
func VerifyJWT(tok string) (models.Claims, error) {
	err := godotenv.Load()
	if err != nil {
		return models.Claims{}, err
	}

	claims := &models.Claims{} // pointer is required here
	token, err := jwt.ParseWithClaims(tok, claims, func(t *jwt.Token) (interface{}, error) {
		// Check signing method
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return models.Claims{}, err
	}

	if !token.Valid {
		return models.Claims{}, errors.New("token provided does not match jwt secret")
	}

	return *claims, nil // dereference before returning
}

func CreateToken(username string) (string, error) {
	err := godotenv.Load()
	if err != nil {
		return "", err
	}

	claims := models.Claims{Username: username, RegisteredClaims: jwt.RegisteredClaims{IssuedAt: jwt.NewNumericDate(time.Now()), ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour))}}

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)

	// we need to sign the token with the secret

	jwt_secret := os.Getenv("JWT_SECRET")

	signedToken, err := token.SignedString([]byte(jwt_secret))
	if err != nil {
		panic("Not able to sign the token")
	}
	fmt.Println("token", signedToken)

	return signedToken, nil
}

func HashBinaryData(data []byte, hash_cost int) ([]byte, error) {

	hashed_password, err := bcrypt.GenerateFromPassword(data, hash_cost)
	if err != nil {
		return nil, err
	}
	return hashed_password, nil
}

func CompareHashAndPassword(data, original_data string) error {
	err := bcrypt.CompareHashAndPassword([]byte(data), []byte(original_data))
	if err != nil {
		return err
	}
	return nil
}
