package auth

import "testing"

func TestCreateJWT(t *testing.T) {
	token, err := CreateJWT([]byte("somesecret"), 2)

	if err != nil {
		t.Errorf("error creating token: %v", err)
	}

	if token == "" {
		t.Error("token should not be empty")
	}
}
