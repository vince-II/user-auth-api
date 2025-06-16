package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/vince-II/auth-post-api/server/helpers"
)

func ConnectToDb(dbUrl string) *pgxpool.Pool {
	// open database
	pool, err := pgxpool.New(context.Background(), dbUrl)
	helpers.CheckError(err)

	return pool
}
