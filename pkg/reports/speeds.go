package reports

import (
	"context"
	"time"

	"github.com/marc-campbell/nicedishy/pkg/persistence"
)

func getFastestDownloadSpeed(ctx context.Context, dishyID string, start time.Time, end time.Time) (float64, error) {
	pg := persistence.MustGetPGSession()

	query := `select max(download_speed) from from dishy_speed where dishy_id = $1 and time > $2 and time < $3`
	rows := pg.QueryRow(ctx, query, dishyID, start, end)

	fastestDownloadSpeed := float64(0)
	if err := rows.Scan(&fastestDownloadSpeed); err != nil {
		return 0, err
	}

	return fastestDownloadSpeed, nil
}

func getFastestUploadSpeed(ctx context.Context, dishyID string, start time.Time, end time.Time) (float64, error) {
	pg := persistence.MustGetPGSession()

	query := `select max(upload_speed) from from dishy_speed where dishy_id = $1 and time > $2 and time < $3`
	rows := pg.QueryRow(ctx, query, dishyID, start, end)

	fastestUploadSpeed := float64(0)
	if err := rows.Scan(&fastestUploadSpeed); err != nil {
		return 0, err
	}

	return fastestUploadSpeed, nil
}

func getAverageDownloadSpeed(ctx context.Context, dishyID string, start time.Time, end time.Time) (float64, error) {
	pg := persistence.MustGetPGSession()

	query := `select avg(download_seed) from from dishy_speed_hourly where dishy_id = $1 and time_start > $2 and time_start < $3`
	rows := pg.QueryRow(ctx, query, dishyID, start, end)

	averageDownloadSeed := float64(0)
	if err := rows.Scan(&averageDownloadSeed); err != nil {
		return 0, err
	}

	return averageDownloadSeed, nil
}

func getAverageUploadSpeed(ctx context.Context, dishyID string, start time.Time, end time.Time) (float64, error) {
	pg := persistence.MustGetPGSession()

	query := `select avg(upload_speed) from from dishy_speed_hourly where dishy_id = $1 and time_start > $2 and time_start < $3`
	rows := pg.QueryRow(ctx, query, dishyID, start, end)

	averageUploadSpeed := float64(0)
	if err := rows.Scan(&averageUploadSpeed); err != nil {
		return 0, err
	}

	return averageUploadSpeed, nil
}
