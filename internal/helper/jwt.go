package helper

import (
	"os"
	"strconv"
	"time"

	"github.com/Dzikuri/shopifyx/internal/model"
	"github.com/golang-jwt/jwt/v5"
)

func JwtSecret() string {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		secret = "secret"
	}

	return secret
}

func JwtExpireAt() int {
	expireAt := os.Getenv("JWT_EXPIRE_AT")
	if expireAt == "" {
		expireAt = "2"
	}

	expire, _ := strconv.Atoi(expireAt)

	return expire
}

type JwtCustomClaims struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	UserId   string `json:"userId"`
	jwt.RegisteredClaims
}

// Valid implements jwt.Claims.
func (j *JwtCustomClaims) Valid() error {
	panic("unimplemented")
}

func JwtGenerateToken(request *model.UserResponse) (string, error) {

	// Generate Claims object
	jwtClaims := JwtCustomClaims{
		Name:     request.Name,
		Username: request.Username,
		UserId:   request.Id.String(),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 2)),
		},
	}

	// Create token with claims
	// token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaims)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(JwtSecret()))
	if err != nil {
		return "", err
	}

	return t, err
}