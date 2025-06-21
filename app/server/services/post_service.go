package services

import (
	"context"
	"errors"

	log "github.com/gofiber/fiber/v2/log"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/vince-II/auth-post-api/connectors"
	"github.com/vince-II/auth-post-api/internal/database"
	"github.com/vince-II/auth-post-api/server/dto"
)

func CreatePost(userID int32, p dto.PostParams, ctx context.Context) (map[string]interface{}, error) {
	pool, err := connectors.ConnectToDb(*connectors.NewDBCredentials(), ctx)
	if err != nil {
		log.Errorf("Failed to connect to database: %v" + err.Error())
		return nil, errors.New("Failed to connect to database")
	}

	exist := doesUserExist(userID, pool)

	if !exist {
		log.Errorf("User doesn't exist: %s")
		return nil, errors.New("Username already taken")
	}

	params := database.CreatePostParams{
		UserID:  userID,
		Content: p.Content,
	}

	result, err := database.New(pool).CreatePost(context.Background(), params)
	if err != nil {
		log.Errorf("Failed to create post %v", err)
		return nil, err
	}
	log.Infof("Post was created successfully %v", result)

	data := map[string]interface{}{
		"post_id": result.ID,
		"content": result.Content,
	}

	return data, nil
}

func doesUserExist(userID int32, pool *pgxpool.Pool) bool {
	exist, err := database.New(pool).UserExists(context.Background(), userID)

	if err != nil {
		log.Errorf("Error checking if username exists: %v", err)
		return false
	}

	return exist
}

func UpdatePost(userID int32, p dto.PostParams, ctx context.Context) (map[string]interface{}, error) {
	pool, err := connectors.ConnectToDb(*connectors.NewDBCredentials(), ctx)
	if err != nil {
		log.Errorf("Failed to connect to database: %v" + err.Error())
		return nil, errors.New("Failed to connect to database")
	}

	exist := doesUserExist(userID, pool)

	if !exist {
		log.Errorf("User doesn't exist: %s")
		return nil, errors.New("Username already taken")
	}

	params := database.UpdatePostParams{
		ID:      int32(p.ID),
		Content: p.Content,
	}

	result, err := database.New(pool).UpdatePost(context.Background(), params)
	if err != nil {
		log.Errorf("Failed to create post %v", err)
		return nil, err
	}

	log.Infof("Post was updated successfully %v", result)

	data := map[string]interface{}{
		"post_id": result.ID,
		"content": result.Content,
	}

	return data, nil
}
