package security

import (
	"time"

	"example.com/gin_forum/config"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(username, email string) (string, error) {
	key := []byte(config.GetConfig())
	tokenDuration := time.Hour * 24
	now := time.Now()
	t := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"user": map[string]string{
				"email":    email,
				"username": username,
			},
			"iat": now.Unix(),                    //表示签发时间
			"exp": now.Add(tokenDuration).Unix(), //表示过期时间
		})
	return t.SignedString(key)

}
