package util

import (
	"errors"
	"os"
	"time"

	"github.com/gofiber/fiber/v2/log"
	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserID   int    `json:"user_id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func GenerateToken(userID int64, username string) (string, error) {
	secretKey := os.Getenv("JWT_SECRET")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"userID":   userID,
			"username": username,
			"exp":      time.Now().Add(time.Hour * 24).Unix(),
			"iat":      time.Now().Unix(), // Issued at
			"iss":      "auth-post-api",
		})

	signedToken, err := token.SignedString(secretKey)

	if err != nil {
		log.Errorf("Failed to signed the token %v", err)
		return "", err
	}
	return signedToken, nil
}

func VerifyToken(tokenStr string) (*Claims, error) {
	secretKey := os.Getenv("JWT_SECRET")

	claims := &Claims{}

	// note: parse vs parsewithclaims
	// parse -> parse the jwt and give generics map claims
	// parsewithclaims -> custom typed + better types access
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("Failed in signing method")
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, errors.New("Invalid Token")
	}

	return claims, nil

}
