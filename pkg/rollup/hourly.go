package rollup

import (
	"context"
	"fmt"
	"time"

	"github.com/marc-campbell/nicedishy/pkg/persistence"
)

func ReindexDataHourly(ctx context.Context, dishyID string, when time.Time) error {
	pg := persistence.MustGetPGSession()

	startHour := when.Truncate(time.Hour)
	endHour := startHour.Add(time.Hour)

	query := `select snr, downlink_throughput_bps, uplink_throughput_bps,
pop_ping_latency_ms, pop_ping_drop_rate, percent_obstructed, seconds_obstructed
from dishy_data where time >= $1 and time < $2 and dishy_id = $3`
	rows, err := pg.Query(ctx, query, startHour, endHour, dishyID)
	if err != nil {
		return fmt.Errorf("error querying for dishy speed: %v", err)
	}
	defer rows.Close()

	metricCount := 0
	totalSNR := 0
	totalDownlinkThroughput := float64(0)
	totalUplinkThroughput := float64(0)
	totalPopPingLatency := float64(0)
	totalPopPingDropRate := float64(0)
	totalPercentObstructed := float64(0)
	totalSecondsObstructed := float64(0)

	for rows.Next() {
		var snr int
		var downlinkThroughputBPS float64
		var uplinkThroughputBPS float64
		var popPingLatencyMS float64
		var popPingDropRate float64
		var percentObstructed float64
		var secondsObstructed float64

		err := rows.Scan(&snr, &downlinkThroughputBPS, &uplinkThroughputBPS, &popPingLatencyMS, &popPingDropRate, &percentObstructed, &secondsObstructed)
		if err != nil {
			return fmt.Errorf("error scanning for dishy data: %v", err)
		}

		totalSNR += snr
		totalDownlinkThroughput += downlinkThroughputBPS
		totalUplinkThroughput += uplinkThroughputBPS
		totalPopPingLatency += popPingLatencyMS
		totalPopPingDropRate += popPingDropRate
		totalPercentObstructed += percentObstructed
		totalSecondsObstructed += secondsObstructed

		metricCount++
	}

	if metricCount == 0 {
		return nil
	}

	avgSNR := float64(totalSNR) / float64(metricCount)
	avgDownlinkThroughput := totalDownlinkThroughput / float64(metricCount)
	avgUplinkThroughput := totalUplinkThroughput / float64(metricCount)
	avgPopPingLatency := totalPopPingLatency / float64(metricCount)
	avgPopPingDropRate := totalPopPingDropRate / float64(metricCount)
	avgPercentObstructed := totalPercentObstructed / float64(metricCount)
	avgSecondsObstructed := totalSecondsObstructed / float64(metricCount)

	query = `insert into dishy_data_hourly (time_start, dishy_id, snr, downlink_throughput_bps, uplink_throughput_bps, pop_ping_latency_ms, pop_ping_drop_rate, percent_obstructed, seconds_obstructed)
values ($1, $2, $3, $4, $5, $6, $7, $8, $9)
on conflict (time_start, dishy_id) do update set
snr = $3, downlink_throughput_bps = $4, uplink_throughput_bps = $5, pop_ping_latency_ms = $6, pop_ping_drop_rate = $7, percent_obstructed = $8, seconds_obstructed = $9`

	_, err = pg.Exec(ctx, query, startHour, dishyID, avgSNR, avgDownlinkThroughput, avgUplinkThroughput, avgPopPingLatency, avgPopPingDropRate, avgPercentObstructed, avgSecondsObstructed)
	if err != nil {
		return fmt.Errorf("error inserting dishy data hourly: %v", err)
	}

	return nil
}

func ReindexSpeedHourly(ctx context.Context, dishyID string, when time.Time) error {
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
