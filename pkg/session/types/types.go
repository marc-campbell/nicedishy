package types

import "time"

type Session struct {
	ID          string    `json:"id"`
	UserID      string    `json:"userId"`
	ExpiresAt   time.Time `json:"expiresAt"`
	AccessToken string    `json:"-"`
}
