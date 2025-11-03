package utils

import (
	"fmt"
	"testing"
)

func TestCreateToken(t *testing.T) {
	data := "jordan"
	token, err := CreateToken(data)
	if err != nil {
		t.Errorf("not a valid test anyways")
	}
	fmt.Println("token: ", string(token))

}

func TestVerifyToken(t *testing.T) {
	data := "jordan"
	token, err := CreateToken(data)
	if err != nil {
		t.Error("not a valid test anyways", err.Error())
	}
	mapClaims, err := VerifyJWT(string(token))
	if err != nil {
		t.Error("not a valid test anyways", err.Error())
	}
	fmt.Println("claims", mapClaims)
	fmt.Println(mapClaims["username"])
}
