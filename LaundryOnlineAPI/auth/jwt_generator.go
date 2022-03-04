package auth

import (
	"LaundryOnlineAPI/config"
	"time"

	"github.com/golang-jwt/jwt"
)

type JWTClaims struct {
	jwt.StandardClaims
	Username string `json:"username"`
	Password string `json:"password"`
}

func GenerateJWT(username string, password string, issuer string, expireDuration time.Duration) (string, error) {
	var config = config.Config

	claims := JWTClaims{
		StandardClaims: jwt.StandardClaims{
			Issuer:    config.APPLICATION_NAME,
			ExpiresAt: time.Now().Add(expireDuration).Unix(),
		},
		Username: username,
		Password: password,
	}

	token := jwt.NewWithClaims(
		config.JWT_SIGNING_METHOD,
		claims,
	)

	signedToken, err := token.SignedString(config.JWT_SIGNATURE_KEY)
	if err != nil {
		return "", err
	}
	return signedToken, nil
}
