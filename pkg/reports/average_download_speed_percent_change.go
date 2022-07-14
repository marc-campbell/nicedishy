package reports

import (
	"context"
	"time"

	"github.com/marc-campbell/nicedishy/pkg/logger"
	"github.com/marc-campbell/nicedishy/pkg/persistence"
)

// getAverageDownloadSpeedPercentChange
// takes the current week start/end
// and determines the average download speed percent change from the previous week
func getAverageDownloadSpeedPercentChange(ctx context.Context, dishyID string, start time.Time, end time.Time) (float64, error) {
	// get the previous week start/end
	prevStart, prevEnd := getPreviousWeekStartEnd(start, end)

	pg := persistence.MustGetPGSession()

	// get the average download speed from this week
	query := `select download_speed from from dishy_speed_hourly where dishy_id = $1 and time_start > $2 and time_start < $3`
	rows, err := pg.Query(ctx, query, dishyID, start, end)
	if err != nil {
		logger.Error(err)
		return 0, err
	}
	defer rows.Close()

	var total float64
	var count int
	for rows.Next() {
		var downloadSpeed float64
		if err := rows.Scan(&downloadSpeed); err != nil {
			logger.Error(err)
			return 0, err
		}

		total += downloadSpeed
		count++
	}

	rows.Close()

	// get last week average download speed
	query = `select download_speed from from dishy_speed_hourly where dishy_id = $1 and time_start > $2 and time_start < $3`
	rows, err = pg.Query(ctx, query, dishyID, prevStart, prevEnd)
	if err != nil {
		logger.Error(err)
		return 0, err
	}
	defer rows.Close()

	var prevTotal float64
	var prevCount int
	for rows.Next() {
		var downloadSpeed float64
		if err := rows.Scan(&downloadSpeed); err != nil {
			logger.Error(err)
			return 0, err
		}

		prevTotal += downloadSpeed
		prevCount++
	}

	if prevTotal == 0 {
		return 0, nil
	}

	return (total - prevTotal) / prevTotal, nil
}
