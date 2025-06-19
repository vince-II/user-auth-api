package connectors

import (
	"context"
	"fmt"

	"os"

	"github.com/gofiber/fiber/v2/log"
	"github.com/jackc/pgx/v5/pgxpool"
)

type DBCredentials struct {
	host     string
	user     string
	password string
	dbName   string
	port     string
}

func NewDBCredentials() *DBCredentials {
	// loads environment variables for database connection
	var (
		host     = os.Getenv("DB_HOST")
		user     = os.Getenv("DB_USER")
		password = os.Getenv("DB_PASSWORD")
		dbName   = os.Getenv("DB_NAME")
		port     = os.Getenv("DB_PORT")
	)
	return &DBCredentials{host: host, user: user, password: password, dbName: dbName, port: port}
}

func ConnectToDb(creds DBCredentials, ctx context.Context) (*pgxpool.Pool, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		creds.host, creds.port, creds.user, creds.password, creds.dbName,
	)

	// open database
	pool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		log.Errorf("Failed to connect to database %v", err)
		return nil, err
	}

	return pool, nil
}
