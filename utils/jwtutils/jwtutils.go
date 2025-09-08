package jwtutils

import (
	"os"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(sub any) (string, error) {
	var claims = jwt.MapClaims{
		"sub": sub,
		"exp": time.Now().Add(time.Minute * 15).Unix(),
	}
	var jwt = jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := jwt.SignedString([]byte(os.Getenv("KEY")))
	if err != nil {
		return "", err
	}
	return t, nil
}

func VerifyToken(c fiber.Ctx) any {
	var jwtdata = c.Locals("user").(*jwt.Token)
	var claims = jwtdata.Claims.(jwt.MapClaims)
	var sub = claims["sub"].(float64)
	var exp = claims["exp"].(float64)
	var data = map[string]any{
		"sub": sub,
		"exp": exp,
	}
	return data
}
