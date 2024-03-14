package security

import "testing"

func TestGenerateJWT(t *testing.T) {
	token, err := GenerateJWT("jack", "jack@gmail.com")
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("token: %v\n", token)
}
