package types

import "time"

type Dishy struct {
	ID           string     `json:"id"`
	CreatedAt    time.Time  `json:"createdAt"`
	LastMetricAt *time.Time `json:"lastMetricAt,omitempty"`
	Name         string     `json:"name"`
}
