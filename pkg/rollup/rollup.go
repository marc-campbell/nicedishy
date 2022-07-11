package rollup

import (
	"context"
	"fmt"
	"time"

	"github.com/marc-campbell/nicedishy/pkg/logger"
	"github.com/marc-campbell/nicedishy/pkg/persistence"
)

func ReindexAll() error {
	pg := persistence.MustGetPGSession()

	query := `select id from dishy`
	rows, err := pg.Query(context.Background(), query)
	if err != nil {
		return fmt.Errorf("error querying for dishy: %v", err)
	}
	defer rows.Close()

	dishyIDs := []string{}
	for rows.Next() {
		var dishyID string
		err := rows.Scan(&dishyID)
		if err != nil {
			return fmt.Errorf("error scanning for dishy: %v", err)
		}

		dishyIDs = append(dishyIDs, dishyID)
	}
	rows.Close()

	for _, dishyID := range dishyIDs {
		logger.Infof("reindexing dishy %s", dishyID)
		min, err := time.Parse(time.RFC3339, "2022-01-01T00:01:00Z")
		if err != nil {
			return fmt.Errorf("error parsing start time: %v", err)
		}

		max := time.Now()

		for min.Before(max) {
			fmt.Printf("reindexing %s from %s to %s\n", dishyID, min, min.Add(time.Hour))
			if err := ReindexHourly(context.Background(), dishyID, min); err != nil {
				return fmt.Errorf("error reindexing hourly: %v", err)
			}

			min = min.Add(time.Hour)

		}

	}

	return nil
}

func ReindexHourly(ctx context.Context, dishyID string, when time.Time) error {
	pg := persistence.MustGetPGSession()

	startHour := when.Truncate(time.Hour)
	endHour := startHour.Add(time.Hour)

	query := `select download_speed, upload_speed from dishy_speed where time >= $1 and time < $2 and dishy_id = $3`
	rows, err := pg.Query(ctx, query, startHour, endHour, dishyID)
	if err != nil {
		return fmt.Errorf("error querying for dishy speed: %v", err)
	}
	defer rows.Close()

	metricCount := 0
	totalDownloadSpeed := float64(0)
	totalUploadSpeed := float64(0)

	for rows.Next() {
		var downloadSpeed float64
		var uploadSpeed float64
		err := rows.Scan(&downloadSpeed, &uploadSpeed)
		if err != nil {
			return fmt.Errorf("error scanning for dishy speed: %v", err)
		}

		metricCount++
		totalDownloadSpeed += downloadSpeed
		totalUploadSpeed += uploadSpeed
	}

	if metricCount == 0 {
		return nil
	}

	averageDownloadSpeed := totalDownloadSpeed / float64(metricCount)
	averageUploadSpeed := totalUploadSpeed / float64(metricCount)

	query = `insert into dishy_speed_hourly (time_start, dishy_id, download_speed, upload_speed) values ($1, $2, $3, $4)
on conflict (time_start, dishy_id) do update set download_speed = $3, upload_speed = $4`

	_, err = pg.Exec(ctx, query, startHour, dishyID, averageDownloadSpeed, averageUploadSpeed)
	if err != nil {
		return fmt.Errorf("error inserting dishy speed: %v", err)
	}

	return nil
}
