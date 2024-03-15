package middleware

import (
	"fmt"
	"log"
	"strings"

	"github.com/Dzikuri/shopifyx/internal/helper"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func JwtCheckTokenUser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Request().Header.Get("Authorization")
		if token == "" {
            return echo.ErrUnauthorized
		}
        
		tokenString := strings.Split(token, "Bearer ")[1]

        // Extract token from header
		t, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
            // Don't forget to validate the alg is what you expect:
            if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
                return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
            }

            return []byte(helper.JwtSecret()), nil
        })
        if err != nil {
            log.Fatal(err)
            return echo.ErrUnauthorized
        }

        if claims, ok := t.Claims.(jwt.MapClaims); ok {
            c.Set("user", claims)
			return next(c)
        } else {
            fmt.Println(err)
            return echo.ErrUnauthorized
        }
	}
}