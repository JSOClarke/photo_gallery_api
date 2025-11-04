package utils

import (
	"fmt"
	"strings"
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
	assert.GreaterOrEqual(t, len(token), 100)
	assert.Equal(t, 3, len(strings.Split(token, ".")))

}

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
	fmt.Println("claims ->", claims)
	assert.Equal(t, "Jordan", claims.Username)
}

func TestHashPassword(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		password  []byte
		hash_cost int
		want      []byte
		wantErr   bool
	}{
		// TODO: Add test cases.
		{
			name:      "valid password",
			password:  []byte("secret123"),
			hash_cost: 10,
			wantErr:   false,
		},
		{
			name:      "too low cost",
			password:  []byte("secret123"),
			hash_cost: 1,
			wantErr:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := HashBinaryData(tt.password, tt.hash_cost)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("HashPassword() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("HashPassword() succeeded unexpectedly")
			}
			// TODO: update the condition below to compare got with tt.want.
			if !tt.wantErr && len(got) == 0 {
				t.Errorf("HashPassword() returned and empty Hash")
			}
		})
	}
}

func TestCompareHashAndPassword(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		data          string
		original_data string
		hash_cost     int
		wantErr       bool
	}{
		// TODO: Add test cases.
		{name: "valid paramters",
			data:      "Jordan",
			hash_cost: 10,
			wantErr:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hashed, err := HashBinaryData([]byte(tt.data), 10)
			if err != nil {
				t.Errorf("HashBinaryData() failed, %s", err.Error())
			}
			gotErr := CompareHashAndPassword(string(hashed), tt.data)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("CompareHashAndPassword() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("CompareHashAndPassword() succeeded unexpectedly")
			}
		})
	}
}
