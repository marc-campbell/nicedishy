package persistence

import (
	"context"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
)

var MetricsDB *pgxpool.Pool

func MustGetMetricsDBSession() *pgxpool.Pool {
	if MetricsDB != nil {
		return MetricsDB
	}

	config, err := pgxpool.ParseConfig(os.Getenv("METRICS_DB_URI"))
	if err != nil {
		panic(err)
	}

	conn, err := pgxpool.ConnectConfig(context.Background(), config)
	if err != nil {
		panic(err)
	}

	MetricsDB = conn
	return MetricsDB
}
