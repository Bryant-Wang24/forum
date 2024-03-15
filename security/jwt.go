package security

import (
	"crypto/rsa"
	"os"
	"time"

	"example.com/gin_forum/config"
	"github.com/golang-jwt/jwt/v5"
)

var privateKey *rsa.PrivateKey
var publicKey *rsa.PublicKey

func init() {
	var err error
	var bytes []byte
	bytes, err = os.ReadFile(config.GetPrivateKeyLocation())
	if err != nil {
		panic(err)
	}

	privateKey, err = jwt.ParseRSAPrivateKeyFromPEM(bytes)
	if err != nil {
		panic(err)
	}

	bytes, err = os.ReadFile(config.GetPublicKeyLocation())
	if err != nil {
		panic(err)
	}
	publicKey, err = jwt.ParseRSAPublicKeyFromPEM(bytes)
	if err != nil {
		panic(err)
	}
}

func GenerateJWT(username, email string) (string, error) {
	// key := []byte(config.GetConfig())
	tokenDuration := time.Hour * 24
	now := time.Now()
	t := jwt.NewWithClaims(jwt.SigningMethodRS512,
		jwt.MapClaims{
			"user": map[string]string{
				"email":    email,
				"username": username,
			},
			"iat": now.Unix(),                    //表示签发时间
			"exp": now.Add(tokenDuration).Unix(), //表示过期时间
		})
	return t.SignedString(privateKey)

}

func VerifyJWT(token string) (*jwt.MapClaims, bool, error) {
	var claim jwt.MapClaims
	claims, err := jwt.ParseWithClaims(token, &claim, func(t *jwt.Token) (interface{}, error) {
		return publicKey, nil
	})
	if err != nil {
		return nil, false, err
	}
	if claims.Valid { //如果token有效，返回true
		return &claim, true, nil
	}
	return nil, false, nil
}
