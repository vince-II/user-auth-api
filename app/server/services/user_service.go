package services

import (
	"context"
	"errors"

	"github.com/vince-II/auth-post-api/connectors"
	"github.com/vince-II/auth-post-api/internal/database"
	"github.com/vince-II/auth-post-api/server/dto"
	"github.com/vince-II/auth-post-api/server/util"

	log "github.com/gofiber/fiber/v2/log"
	"github.com/jackc/pgx/v5/pgxpool"
)

func RegisterUser(user dto.RegisterUser, ctx context.Context) (map[string]interface{}, error) {
	pool, err := connectors.ConnectToDb(*connectors.NewDBCredentials(), ctx)

	if err != nil {
		log.Errorf("Failed to connect to database: %v" + err.Error())
		return nil, errors.New("Failed to connect to database")
	}

	alreadyExist := doesUsernameExist(user.Username, pool)

	if alreadyExist {
		log.Errorf("Username already taken: %s", user.Username)
		return nil, errors.New("Username already taken")
	}

	hashedPassword, err := util.HashPassword(user.Password)
	if err != nil {
		log.Errorf(err.Error())
		return nil, err
	}

	params := database.CreateUserParams{
		Username:  user.Username,
		Password:  hashedPassword,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}

	data, err := database.New(pool).CreateUser(context.Background(), params)

	if err != nil {
		log.Errorf(err.Error())
		return nil, err
	}

	log.Infof("User registered successfully %v", data)
	result := map[string]interface{}{
		"id":         data.ID,
		"username":   data.Username,
		"first_name": data.FirstName,
		"last_name":  data.LastName,
	}

	return result, nil

}

func doesUsernameExist(username string, pool *pgxpool.Pool) bool {
	exist, err := database.New(pool).UsernameExists(context.Background(), username)

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
