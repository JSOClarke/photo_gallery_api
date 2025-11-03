package utils

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateToken(t *testing.T) {
	data := "jordan"
	token, err := CreateToken(data)
	if err != nil {
		t.Errorf("not a valid test anyways")
	}
	fmt.Println("token: ", string(token))

}

// Tests for claims username
func TestVerifyToken(t *testing.T) {
	data := "Jordan"
	token, err := CreateToken(data)
	if err != nil {
		t.Error("not a valid test anyways", err.Error())
	}
	claims, err := VerifyJWT(string(token))
	if err != nil {
		t.Error("not a valid test anyways", err.Error())
	}
	assert.Equal(t, "Jordan", claims.Username)
}
