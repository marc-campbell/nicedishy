package types

import "time"

type Dishy struct {
	ID           string     `json:"id"`
	CreatedAt    time.Time  `json:"createdAt"`
	LastMetricAt *time.Time `json:"lastMetricAt,omitempty"`
	Name         string     `json:"name"`
}

type DishyStat struct {
	State                 string  `json:"state"`
	SNR                   int     `json:"snr"`
	DownlinkThroughputBps float64 `json:"downloadThroughputBps"`
	UplinkThroughputBps   float64 `json:"uplinkThroughputBps"`
	PopPingLatencyMs      float64 `json:"popPingLatencyMs"`
	PopPingDropRate       float64 `json:"popPingDropRate"`
	PercentObstructed     float64 `json:"percentObstructed"`
	ObstructedSeconds     float64 `json:"obstructedSeconds"`
}

type DishyWithStats struct {
	Dishy
	Latest      *DishyStat               `json:"latest"`
	RecentStats map[time.Time]*DishyStat `json:"recent"`
}
