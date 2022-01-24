package pgstore

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/marc-campbell/nicedishy/pkg/persistence"
)

func (s PGStore) GetConnectedDishyCount(ctx context.Context) (int, error) {
	pg := persistence.MustGetPGSession()

	query := `select count(1) from dishy where last_metric_at > $1`
	row := pg.QueryRow(ctx, query, time.Now().Add(-24*time.Hour))

	count := 0
	if err := row.Scan(&count); err != nil {
		return 0, fmt.Errorf("scan count: %w", err)
	}

	return count, nil
}

func (s PGStore) GetTotalDishyCount(ctx context.Context) (int, error) {
	pg := persistence.MustGetPGSession()

	query := `select count(1) from dishy where last_metric_at is not null`
	row := pg.QueryRow(ctx, query)

	count := 0
	if err := row.Scan(&count); err != nil {
		return 0, fmt.Errorf("scan count: %w", err)
	}

	return count, nil
}

func (s PGStore) GetNewDishyCount(ctx context.Context) (int, error) {
	pg := persistence.MustGetPGSession()

	query := `select count(1) from dishy where last_metric_at is not null `
	row := pg.QueryRow(ctx, query)

	count := 0
	if err := row.Scan(&count); err != nil {
		return 0, fmt.Errorf("scan count: %w", err)
	}

	return count, nil
}

func (s PGStore) GetHighestDownloadSpeed(ctx context.Context) (float64, error) {
	pg := persistence.MustGetPGSession()

	query := `select max(download_speed) from dishy_speed`
	row := pg.QueryRow(ctx, query)

	var speed sql.NullFloat64
	if err := row.Scan(&speed); err != nil {
		return 0, fmt.Errorf("scan download speed: %w", err)
	}

	if !speed.Valid {
		return 0.0, nil
	}

	return speed.Float64, nil
}

func (s PGStore) GetAverageDownloadSpeed(ctx context.Context) (float64, error) {
	pg := persistence.MustGetPGSession()

	query := `select avg(download_speed) from dishy_speed`
	row := pg.QueryRow(ctx, query)

	var speed sql.NullFloat64
	if err := row.Scan(&speed); err != nil {
		return 0, fmt.Errorf("scan upload speed: %w", err)
	}

	if !speed.Valid {
		return 0.0, nil
	}

	return speed.Float64, nil
}

func (s PGStore) GetLowestPingTime(ctx context.Context) (float64, error) {
	pg := persistence.MustGetPGSession()

	query := `select min(pop_ping_latency_ms) from dishy_data`
	row := pg.QueryRow(ctx, query)

	var min sql.NullFloat64
	if err := row.Scan(&min); err != nil {
		return 0.0, fmt.Errorf("scan min: %w", err)
	}

	if !min.Valid {
		return 0.0, nil
	}

	return min.Float64, nil
}

func (s PGStore) GetAveragePingTime(ctx context.Context) (float64, error) {
	pg := persistence.MustGetPGSession()

	query := `select avg(pop_ping_latency_ms) from dishy_data`
	row := pg.QueryRow(ctx, query)

	var avg sql.NullFloat64
	if err := row.Scan(&avg); err != nil {
		return 0.0, fmt.Errorf("scan avg: %w", err)
	}

	if !avg.Valid {
		return 0.0, nil
	}

	return avg.Float64, nil
}
