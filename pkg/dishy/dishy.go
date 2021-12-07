package dishy

import (
	"context"
	"fmt"
	"time"

	"github.com/marc-campbell/nicedishy/pkg/dishy/types"
	"github.com/marc-campbell/nicedishy/pkg/persistence"
)

// GetLatestStats will query timescale for the current stats
// for this dishy id
func GetLatestStats(id string) (*types.DishyStat, error) {
	metricsDB := persistence.MustGetMetricsDBSession()
	query := `select state, snr, downlink_throughput_bps, uplink_throughput_bps, pop_ping_latency_ms, pop_ping_drop_rate, percent_obstructed, seconds_obstructed from dishy_data where dishy_id = $1 order by time desc limit 1`
	row := metricsDB.QueryRow(context.Background(), query, id)

	stats := types.DishyStat{}
	if err := row.Scan(&stats.State, &stats.SNR, &stats.DownlinkThroughputBps, &stats.UplinkThroughputBps, &stats.PopPingLatencyMs, &stats.PopPingDropRate, &stats.PercentObstructed, &stats.ObstructedSeconds); err != nil {
		return nil, fmt.Errorf("error scanning stats: %w", err)
	}

	return &stats, nil
}

func GetRecentStats(id string) (map[time.Time]*types.DishyStat, error) {
	metricsDB := persistence.MustGetMetricsDBSession()
	query := `select time, state, snr, downlink_throughput_bps, uplink_throughput_bps, pop_ping_latency_ms, pop_ping_drop_rate, percent_obstructed, seconds_obstructed from dishy_data where dishy_id = $1 order by time desc limit 10`
	rows, err := metricsDB.Query(context.Background(), query, id)
	if err != nil {
		return nil, fmt.Errorf("error querying: %w", err)
	}

	recent := map[time.Time]*types.DishyStat{}
	for rows.Next() {
		stats := types.DishyStat{}
		when := time.Time{}
		if err := rows.Scan(&when, &stats.State, &stats.SNR, &stats.DownlinkThroughputBps, &stats.UplinkThroughputBps, &stats.PopPingLatencyMs, &stats.PopPingDropRate, &stats.PercentObstructed, &stats.ObstructedSeconds); err != nil {
			return nil, fmt.Errorf("error scanning stats: %w", err)
		}

		recent[when] = &stats
	}

	return recent, nil
}
