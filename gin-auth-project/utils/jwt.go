package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(userID uint, role string) (string, error) {
	secret := os.Getenv("JWT_SECRET")

	claims := jwt.MapClaims{
		"user_id": userID,
		"role":    role,
		"exp":     time.Now().Add(time.Hour * 24).Unix(), // expires in 1 day
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//fmt.Println(claims)
	return token.SignedString([]byte(secret))
}
