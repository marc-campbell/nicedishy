package reports

import (
	"context"
	"time"

	"github.com/marc-campbell/nicedishy/pkg/persistence"
)

func getFastestDownloadSpeedTimeBlock(ctx context.Context, dishyID string, start time.Time, end time.Time) (*time.Time, *time.Time, float64, error) {
	pg := persistence.MustGetPGSession()

	query := `select max(download_speed) from from dishy_speed where dishy_id = $1 and time > $2 and time < $3`
	rows := pg.QueryRow(ctx, query, dishyID, start, end)

	fastestDownloadSpeed := float64(0)
	if err := rows.Scan(&fastestDownloadSpeed); err != nil {
		return nil, nil, 0, err
	}

	return nil, nil, fastestDownloadSpeed, nil
}

func getFastestUploadSpeedTimeBlock(ctx context.Context, dishyID string, start time.Time, end time.Time) (*time.Time, *time.Time, float64, error) {
	pg := persistence.MustGetPGSession()

	query := `select max(upload_speed) from from dishy_speed where dishy_id = $1 and time > $2 and time < $3`
	rows := pg.QueryRow(ctx, query, dishyID, start, end)

	uploadDownloadSpeed := float64(0)
	if err := rows.Scan(&uploadDownloadSpeed); err != nil {
		return nil, nil, 0, err
	}

	return nil, nil, uploadDownloadSpeed, nil
}
