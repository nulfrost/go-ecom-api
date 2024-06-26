package auth

import "testing"

func TestHashPassword(t *testing.T) {
	hash, err := HashPassword("password")

	if err != nil {
		t.Errorf("error hashing password: %v", err)
	}

	if hash == "" {
		t.Errorf("password has should not be empty")
	}

	if hash == "password" {
		t.Error("password was not hashed correctly")
	}

}

func TestComparePassword(t *testing.T) {
	hash, err := HashPassword("password")

	if err != nil {
		t.Errorf("error hashing password: %v", err)
	}

	if !ComparePassword(hash, []byte("password")) {
		t.Errorf("password does not match hash")
	}

	if ComparePassword(hash, []byte("notpassword")) {
		t.Errorf("password should not match hash")
	}
}
