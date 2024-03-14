package helper

import (
	"os"
	"strconv"

	"github.com/golang-jwt/jwt/v5"
)

func JwtSecret() string {
	secret := os.Getenv("JWT_SECRET")
	if secret == ""  {
		secret = "secret"
	}

	return secret
}

func JwtExpireAt() int {
	expireAt := os.Getenv("JWT_EXPIRE_AT")
	if expireAt == ""  {
		expireAt = "2"
	}

	expire, _ := strconv.Atoi(expireAt)

	return expire
}

type JwtCustomClaims struct {
	Name   string `json:"name"`
	Username string `json:"username"`
	UserId string `json:"userId"`
	jwt.RegisteredClaims
}
