package types

import "time"

type Dishy struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	Name      string    `json:"name"`
}
