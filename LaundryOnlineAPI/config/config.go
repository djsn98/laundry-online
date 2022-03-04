package config

import (
	"time"

	"github.com/golang-jwt/jwt"
)

var LOGIN_EXPIRATION_DURATION = time.Duration(1) * time.Hour
var JWT_SIGNING_METHOD = jwt.SigningMethodHS256
var JWT_SIGNATURE_KEY = []byte("secret-key")

type ConfigModel struct {
	APPLICATION_NAME          string
	HOST                      string
	PORT                      string
	DBNAME                    string
	DBUSER                    string
	DBHOST                    string
	DBPORT                    string
	DBPASS                    string
	LOGIN_EXPIRATION_DURATION time.Duration
	JWT_SIGNING_METHOD        *jwt.SigningMethodHMAC
	JWT_SIGNATURE_KEY         []byte
}

var Config ConfigModel
