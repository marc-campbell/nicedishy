package types

import "time"

type Dishy struct {
	ID             string     `json:"id"`
	CreatedAt      time.Time  `json:"createdAt"`
	LastMetricAt   *time.Time `json:"lastMetricAt,omitempty"`
	LastGeocheckAt *time.Time `json:"-"`
	Name           string     `json:"name"`
}

type DishyStat struct {
	SNR                   int     `json:"snr"`
	DownlinkThroughputBps float64 `json:"downloadThroughputBps"`
	UplinkThroughputBps   float64 `json:"uplinkThroughputBps"`
	PopPingLatencyMs      float64 `json:"popPingLatencyMs"`
	PopPingDropRate       float64 `json:"popPingDropRate"`
	PercentObstructed     float64 `json:"percentObstructed"`
	ObstructedSeconds     float64 `json:"obstructedSeconds"`
}

type DishySpeed struct {
	DownloadSpeed float64 `json:"downloadSpeed"`
	UploadSpeed   float64 `json:"uploadSpeed"`
}

type DishyWithStats struct {
	Dishy
	LatestStats  *DishyStat  `json:"latestStats"`
	LatestSpeeds *DishySpeed `json:"latestSpeeds"`
}

// https://ipwhois.io/documentation
type GeoCheck struct {
	IPAddress string `json:"ip"`
	Continent string `json:"continent"`
	Country   string `json:"country"`
	Region    string `json:"region"`

	City string `json:"city"`

	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`

	Org string `json:"org"`

	Timezone GeoCheckTimezone `json:"timezone"`
}

type GeoCheckTimezone struct {
	ID     string `json:"id"`
	Abbr   string `json:"abbr"`
	UTC    string `json:"utc"`
	Offset int    `json:"offset"`
}
