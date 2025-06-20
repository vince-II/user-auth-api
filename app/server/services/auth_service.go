package services

import (
	"context"
	"errors"

	log "github.com/gofiber/fiber/v2/log"
	"github.com/vince-II/auth-post-api/connectors"
	"github.com/vince-II/auth-post-api/internal/database"
	"github.com/vince-II/auth-post-api/server/dto"
	"github.com/vince-II/auth-post-api/server/util"
)

func LoginUser(user dto.LoginUser, ctx context.Context) (map[string]interface{}, error) {
	pool, err := connectors.ConnectToDb(*connectors.NewDBCredentials(), ctx)
	if err != nil {
		log.Errorf("Failed to connect to database: %v" + err.Error())
		return nil, errors.New("Failed to connect to database")
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
		log.Errorf("Invalid password for username: %s", user.Username)
		return nil, errors.New("Invalid credentials")
	}

	if err := database.New(pool).UpdateLastLogin(ctx, data.ID); err != nil {
		log.Errorf("Could not update last_login for user %s: %v", data.Username, err)
		return nil, errors.New("Failed to update login time")

	}

	signedToken, _ := util.GenerateToken(int64(data.ID), data.Username)

	if signedToken == "" {
		return nil, errors.New("Failed to generate JWT Token")
	}

	result := map[string]interface{}{
		"token": signedToken,
		"user": &dto.User{
			ID:       int(data.ID),
			Username: data.Username,
		},
	}

	log.Infof("Login Result %v", result)
	return result, nil
}

func LogoutUser(id int32, ctx context.Context) error {
	pool, err := connectors.ConnectToDb(*connectors.NewDBCredentials(), ctx)
	if err != nil {
		log.Errorf("Failed to connect to database: %v" + err.Error())
		return errors.New("Failed to connect to database")
	}

	if err := database.New(pool).UpdateLastLogout(ctx, id); err != nil {
		log.Errorf("Failed to update the logout time: %v", err)
		return errors.New(("Failed to update the logout time"))
	}

	return nil
}
