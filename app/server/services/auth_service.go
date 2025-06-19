package services

import (
	"context"
	"errors"
	"os"
	"time"

	log "github.com/gofiber/fiber/v2/log"
	"github.com/golang-jwt/jwt/v5"
	"github.com/vince-II/auth-post-api/connectors"
	"github.com/vince-II/auth-post-api/internal/database"
	"github.com/vince-II/auth-post-api/server/dto"
	"github.com/vince-II/auth-post-api/server/util"
)

func LoginUser(user dto.LoginUser, ctx context.Context) (map[string]interface{}, error) {
	pool, err := connectors.ConnectToDb(*connectors.NewDBCredentials(), ctx)

	if err != nil {
		panic("Failed to connect to database: " + err.Error())
	}
	exist := doesUsernameExist(user.Username, pool)

	if !exist {
		log.Infof("Username not found: %s", user.Username)
		return nil, errors.New("Username not found")
	}

	data, err := database.New(pool).GetUserByUsername(context.Background(), user.Username)
	if err != nil {
		log.Errorf(err.Error())
		return nil, err
	}

	hashedPassword := data.Password

	if !util.CheckPasswordHash(user.Password, hashedPassword) {
		log.Infof("Invalid password for username: %s", user.Username)
		return nil, errors.New("Invalid credentials")
	}

	signedToken, _ := GenerateJWT(int64(data.ID), data.Username)

	if signedToken == "" {
		return nil, errors.New("Failed to generate JWT Token")
	}

	result := map[string]interface{}{
		"token": signedToken,
		"user": &dto.User{
			ID:        int(data.ID),
			Username:  data.Username,
			FirstName: data.FirstName,
			LastName:  data.LastName,
		},
	}

	log.Infof("Login Result %v", result)
	return result, nil
}

func GenerateJWT(userID int64, username string) (string, error) {
	secret := os.Getenv("JWT_SECRET")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"userID":   userID,
			"username": username,
			"exp":      time.Now().Add(time.Hour * 24).Unix(),
		})

	signedToken, err := token.SignedString(secret)

	if err != nil {
		log.Errorf("Failed to signed the token %v", err)
		return "", err
	}
	return signedToken, nil
}
