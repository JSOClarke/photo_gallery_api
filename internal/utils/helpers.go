package utils

import (
	"errors"
	"fmt"
	"os"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

// Create a model and tests for this.
func VerifyJWT(tok string) (jwt.MapClaims, error) {

	err := godotenv.Load()
	if err != nil {
		return nil, err
	}
	token, err := jwt.Parse(tok, func(t *jwt.Token) (interface{}, error) {
		// check signing method
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}

		// return the key to use for verification
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("Token provided doesnt not match jwt secret")
	}
	claims := token.Claims.(jwt.MapClaims)
	return claims, nil

}

func CreateToken(username string) (string, error) {
	err := godotenv.Load()
	if err != nil {
		return "", err
	}

	claims := jwt.MapClaims{"username": username}

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
