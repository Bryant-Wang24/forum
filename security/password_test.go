package security

import "testing"

func TestHashPassword(t *testing.T) {
	hashPassword, err := HashPassword("secret123")
	if err != nil {
		t.Errorf("hash password failed, err: %v", err)
		return
	}

	t.Logf("hash password: %v\n", hashPassword)

	check := CheckPassword("secret123", hashPassword)
	if !check {
		t.Errorf("check password failed")
		return
	}
	t.Logf("check: %v\n", check)
}
