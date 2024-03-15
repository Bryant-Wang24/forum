package security

import (
	"encoding/base64"

	"golang.org/x/crypto/bcrypt"
)

const salt = "sadfxcvzcxv"

/* 对密码进行加密 */
func HashPassword(password string) (string, error) {
	password += salt
	bcryptedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(bcryptedPassword), nil
}

/* 对密码进行校验 */
func CheckPassword(plain, hash string) bool {
	plain += salt
	bcryptedPassword, err := base64.StdEncoding.DecodeString(hash)
	if err != nil {
		return false
	}
	err = bcrypt.CompareHashAndPassword(bcryptedPassword, []byte(plain))
	return err == nil
}
