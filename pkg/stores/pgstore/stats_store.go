package pgstore

import (
	"context"
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

	query := `select max(download_speed) from dishy_data`
	row := pg.QueryRow(ctx, query)

	speed := 0.0
	if err := row.Scan(&speed); err != nil {
		return 0, fmt.Errorf("scan speed: %w", err)
	}

	return speed, nil
}

func (s PGStore) GetAverageDownloadSpeed(ctx context.Context) (float64, error) {
	pg := persistence.MustGetPGSession()

	query := `select avg(download_speed) from dishy_data`
	row := pg.QueryRow(ctx, query)

	speed := 0.0
	if err := row.Scan(&speed); err != nil {
		return 0, fmt.Errorf("scan speed: %w", err)
	}

	return speed, nil
}

func (s PGStore) GetLowestPingTime(ctx context.Context) (float64, error) {
	pg := persistence.MustGetPGSession()

	query := `select min(pop_ping_latency_ms) from dishy_data`
	row := pg.QueryRow(ctx, query)

	min := 0.0
	if err := row.Scan(&min); err != nil {
		return 0.0, fmt.Errorf("scan min: %w", err)
	}

	return min, nil
}

func (s PGStore) GetAveragePingTime(ctx context.Context) (float64, error) {
	pg := persistence.MustGetPGSession()

	query := `select avg(pop_ping_latency_ms) from dishy_data`
	row := pg.QueryRow(ctx, query)

	avg := 0.0
	if err := row.Scan(&avg); err != nil {
		return 0.0, fmt.Errorf("scan avg: %w", err)
	}

	return avg, nil
}
