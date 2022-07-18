package reports

import (
	"context"
	"time"

	"github.com/marc-campbell/nicedishy/pkg/logger"
	"github.com/marc-campbell/nicedishy/pkg/persistence"
)

func getNumberFirmwareUpdates(ctx context.Context, dishyID string, start time.Time, end time.Time) (int, error) {
	pg := persistence.MustGetPGSession()

	query := `select distinct(dishy_data) from dishy_data where dishy_id = $1 and time_start > $2 and time_start < $3`
	rows, err := pg.Query(ctx, query, dishyID, start, end)
	if err != nil {
		logger.Error(err)
		return 0, err
	}
	defer rows.Close()

	var count int
	for rows.Next() {
		count++
	}

	// count - 1 because it's updates, not total versions
	return count - 1, nil
}
