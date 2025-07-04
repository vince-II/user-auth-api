package services

import (
	"context"
	"errors"
	"fmt"

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

func UpdatePost(postID int32, userID int32, p dto.PostParams, ctx context.Context) (map[string]interface{}, error) {
	pool, err := connectors.ConnectToDb(*connectors.NewDBCredentials(), ctx)
	if err != nil {
		log.Errorf("Failed to connect to database: %v" + err.Error())
		return nil, errors.New("Failed to connect to database")
	}

	exist := doesUserExist(userID, pool)

	if !exist {
		log.Errorf("User doesn't exist: %s")
		return nil, errors.New("User doesn't exist")
	}

	params := database.UpdatePostParams{
		ID:      postID,
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

func DeletePost(postID int32, userID int32, ctx context.Context) error {
	pool, err := connectors.ConnectToDb(*connectors.NewDBCredentials(), ctx)
	if err != nil {
		log.Errorf("Failed to connect to database: %v" + err.Error())
		return errors.New("Failed to connect to database")
	}

	exist := doesUserExist(userID, pool)

	if !exist {
		log.Errorf("User doesn't exist: %s")
		return errors.New("User doesn't exist")
	}

	if err := database.New(pool).DeletePost(context.Background(), postID); err != nil {
		log.Errorf("Failed to delete post %v", err)
		return err
	}

	log.Infof("Post was deleted successfully")

	return nil
}

func GetPost(postID int32, userID int32, ctx context.Context) (map[string]interface{}, error) {
	pool, err := connectors.ConnectToDb(*connectors.NewDBCredentials(), ctx)
	if err != nil {
		log.Errorf("Failed to connect to database: %v" + err.Error())
		return nil, errors.New("Failed to connect to database")
	}

	exist := doesUserExist(userID, pool)

	if !exist {
		log.Errorf("User doesn't exist: %s")
		return nil, errors.New("User doesn't exist")
	}

	result, err := database.New(pool).GetPost(context.Background(), postID)

	if err != nil {
		log.Errorf("Failed to delete post %v", err)
		return nil, fmt.Errorf("%s %v", errors.New("Failed to connect to database"), err)
	}

	log.Infof("Post was read successfully")
	data := map[string]interface{}{
		"post_id": result.ID,
		"content": result.Content,
	}

	return data, nil
}
