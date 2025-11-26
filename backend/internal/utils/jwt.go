package utils

import (
	"eCommerce/internal/config"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var conf = config.LoadEnv()
var secret = conf.JWTSecret

func GenerateToken(userID uint, role string, email string) (string, error) {
	// conf := config.LoadEnv()
	// secret := conf.JWTsecret
	// secret := os.Getenv("JWT_sECRET")
	//fmt.Println(secret, 18)
	claims := jwt.MapClaims{
		"user_id": userID,
		"role":    role,
		"email":   email,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//fmt.Println(claims, 22)
	return token.SignedString([]byte(secret))
}

// func ValidateToken(tokenString string){

// }

// func ValidateToken(tokenString string) (*jwt.Token, jwt.MapClaims, error) {
// 	if secret == "" {
// 		return nil, nil, errors.New("JWT_sECRET is not set", secret)
// 	}

// 	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
// 		// Verify signing method
// 		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
// 			return nil, errors.New("unexpected signing method")
// 		}
// 		return []byte(secret), nil
// 	})

// 	if err != nil {
// 		return nil, nil, err
// 	}

// 	claims, ok := token.Claims.(jwt.MapClaims)
// 	if !ok || !token.Valid {
// 		return nil, nil, errors.New("invalid token")
// 	}

//		return token, claims, nil
//	}
func ValidateToken(tokenString string) (jwt.MapClaims, error) {
	//fmt.Println(secret, 58)
	if secret == "" {
		return nil, errors.New("JWT_sECRET is not set")
	}

	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		// Verify signing method
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(secret), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}
