package utils

import (
	"errors"
	"fmt"
	"os"
	"photogallery/internal/models"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
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
	fmt.Println("jwt_secret", jwt_secret)

	signedToken, err := token.SignedString([]byte(jwt_secret))
	if err != nil {
		panic("Not able to sign the token")
	}
	return signedToken, nil
}
