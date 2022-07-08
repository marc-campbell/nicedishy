package dishy

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/jackc/pgx/v4"
	"github.com/marc-campbell/nicedishy/pkg/dishy/types"
	"github.com/marc-campbell/nicedishy/pkg/persistence"
)

// GetLatestStats will query timescale for the current stats
func GetLatestStats(id string) (*types.DishyStat, error) {
	metricsDB := persistence.MustGetMetricsDBSession()

	query := `select
snr, downlink_throughput_bps, uplink_throughput_bps, pop_ping_latency_ms,
pop_ping_drop_rate, percent_obstructed, seconds_obstructed
from dishy_data
where dishy_id = $1
and downlink_throughput_bps is not null
order by time desc limit 1`
	row := metricsDB.QueryRow(context.Background(), query, id)

	stats := types.DishyStat{}
	if err := row.Scan(&stats.SNR, &stats.DownlinkThroughputBps, &stats.UplinkThroughputBps, &stats.PopPingLatencyMs, &stats.PopPingDropRate,
		&stats.PercentObstructed, &stats.ObstructedSeconds); err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}

		return nil, fmt.Errorf("error scanning stats: %w", err)
	}

	return &stats, nil
}

func GetLatestSpeeds(id string) (*types.DishySpeed, error) {
	metricsDB := persistence.MustGetMetricsDBSession()

	query := `select download_speed, upload_speed from dishy_speed where dishy_id = $1 and download_speed is not null
order by time desc limit 1`
	row := metricsDB.QueryRow(context.Background(), query, id)

	speeds := types.DishySpeed{}
	if err := row.Scan(&speeds.DownloadSpeed, &speeds.UploadSpeed); err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}

		return nil, fmt.Errorf("error scanning speed: %w", err)
	}

	return &speeds, nil
}

func GetRecentStats(id string) (map[time.Time]*types.DishyStat, map[time.Time]*types.DishySpeed, error) {
	metricsDB := persistence.MustGetMetricsDBSession()
	query := `select
time, snr, downlink_throughput_bps, uplink_throughput_bps, pop_ping_latency_ms, pop_ping_drop_rate,
percent_obstructed, seconds_obstructed
from dishy_data
where dishy_id = $1
and downlink_throughput_bps is not null
order by time desc limit 10`
	rows, err := metricsDB.Query(context.Background(), query, id)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil, nil
		}

		return nil, nil, fmt.Errorf("error querying: %w", err)
	}

	stats := map[time.Time]*types.DishyStat{}
	for rows.Next() {
		stat := types.DishyStat{}
		when := time.Time{}
		if err := rows.Scan(&when, &stat.SNR, &stat.DownlinkThroughputBps, &stat.UplinkThroughputBps, &stat.PopPingLatencyMs,
			&stat.PopPingDropRate, &stat.PercentObstructed, &stat.ObstructedSeconds); err != nil {
			return nil, nil, fmt.Errorf("error scanning stats: %w", err)
		}

		stats[when] = &stat
	}

	query = `select
time, download_speed, upload_speed
from dishy_speed
where dishy_id = $1
and download_speed is not null
order by time desc limit 10`
	rows, err = metricsDB.Query(context.Background(), query, id)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil, nil
		}

		return nil, nil, fmt.Errorf("error querying: %w", err)
	}
	speeds := map[time.Time]*types.DishySpeed{}
	for rows.Next() {
		speed := types.DishySpeed{}
		when := time.Time{}

		if err := rows.Scan(&when, &speed.DownloadSpeed, &speed.UploadSpeed); err != nil {
			return nil, nil, fmt.Errorf("error scanning speed: %w", err)
		}

		speeds[when] = &speed
	}

	return stats, speeds, nil
}

// Geocheck will look up the ip address and make sure it looks to be
// a starlink address
func Geocheck(id string, ipAddress string) (*types.GeoCheck, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("https://ipwhois.app/json/%s", ipAddress), nil)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read body: %w", err)
	}

	geoCheck := types.GeoCheck{}
	if err := json.Unmarshal(b, &geoCheck); err != nil {
		return nil, fmt.Errorf("unmarshal json: %w, body was %s", err, b)
	}

	return &geoCheck, nil
}
