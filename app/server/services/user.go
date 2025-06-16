package services

import (
	"context"

	"github.com/vince-II/auth-post-api/internal/sqlc"
	"github.com/vince-II/auth-post-api/server/dto"
	"github.com/vince-II/auth-post-api/server/util"
)

type UserService interface {
	RegisterUser(ctx context.Context) (map[string]interface{}, error)
}

func RegisterUser(user dto.RegisterUser, conn *sqlc.Queries) (map[string]interface{}, error) {
	hashedPassword, err := util.HashPassword(user.Password)
	if err != nil {
		util.LogError(err)
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
		util.LogError(err)
		return nil, err
	}

	util.LogInfo("User registered successfully", data)
	result := map[string]interface{}{
		"id":        data.ID,
		"username":  data.Username,
		"firstName": data.FirstName,
		"lastName":  data.LastName,
	}

	return result, nil

}
