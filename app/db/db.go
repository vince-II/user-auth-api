package database

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

var dbIns *pgxpool.Pool

func Config(ctx context.Context) *pgxpool.Pool {

	if dbIns == nil {
		dbUrl := os.Getenv("DATABASE_URL")

		dbIns = ConnectToDb(dbUrl)
	}

	return dbIns
}
