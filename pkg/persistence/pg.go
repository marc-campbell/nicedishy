package persistence

import (
	"context"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
)

var DB *pgxpool.Pool

func MustGetPGSession() *pgxpool.Pool {
	if DB != nil {
		return DB
	}

	config, err := pgxpool.ParseConfig(os.Getenv("POSTGRES_URI"))
	if err != nil {
		panic(err)
	}

	conn, err := pgxpool.ConnectConfig(context.Background(), config)
	if err != nil {
		panic(err)
	}

	DB = conn
	return DB
}
