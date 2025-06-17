package services

import (
	"context"
	"errors"

	log "github.com/gofiber/fiber/v2/log"
	"github.com/vince-II/auth-post-api/internal/sqlc"
	"github.com/vince-II/auth-post-api/server/dto"
	"github.com/vince-II/auth-post-api/server/util"
)

type UserService interface {
	RegisterUser(ctx context.Context) (map[string]interface{}, error)
}

func RegisterUser(user dto.RegisterUser, conn *sqlc.Queries) (map[string]interface{}, error) {

	isTaken := isUsernameTaken(user.Username, conn)

	if isTaken {
		log.Infof("Username already taken: %s", user.Username)
		return nil, errors.New("Username already taken")
	}

	hashedPassword, err := util.HashPassword(user.Password)
	if err != nil {
		log.Errorf(err.Error())
		return nil, err
	}

	params := sqlc.CreateUserParams{
		Username:  user.Username,
		Password:  hashedPassword,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}

	data, err := conn.CreateUser(context.Background(), params)

	if err != nil {
		log.Errorf(err.Error())
		return nil, err
	}

	log.Infof("User registered successfully", data)
	result := map[string]interface{}{
		"id":        data.ID,
		"username":  data.Username,
		"firstName": data.FirstName,
		"lastName":  data.LastName,
	}

	return result, nil

}

func isUsernameTaken(username string, conn *sqlc.Queries) bool {
	exist, err := conn.UsernameExists(context.Background(), username)

	if err != nil {
		log.Errorf("Error checking if username exists: %v", err)
		return false
	}

	if exist {
		log.Infof("Username already exists %v", username)
		return true
	}

	return false
}
