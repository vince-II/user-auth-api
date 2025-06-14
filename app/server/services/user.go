package services

import (
	"context"

	"github.com/vince-II/auth-post-api/internal/sqlc"
	"github.com/vince-II/auth-post-api/server/models"
	"github.com/vince-II/auth-post-api/server/util"
)

type UserService interface {
	RegisterUser(ctx context.Context) (models.RegisterUser, error)
}

func RegisterUser(user models.RegisterUser, conn *sqlc.Queries) error {
	hashedPassword, err := util.HashPassword(user.Password)
	if err != nil {
		util.LogError(err)
		return nil
	}
	// create a user data
	params := sqlc.CreateUserParams{
		Username:  user.Username,
		Password:  hashedPassword,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}

	data, err := conn.CreateUser(context.Background(), params)
	// create an instance of the database connection

	// create a user login data

	if err != nil {
		util.LogError(err)
		return err
	}
	// param, err = json.Marshal(data)
	util.LogInfo("User registered successfully", data)
	return nil

}
