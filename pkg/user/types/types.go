package types

import (
	"time"
)

type User struct {
	ID           string     `json:"id"`
	EmailAddress string     `json:"emailAddress"`
	AvatarURL    string     `json:"avatarURL"`
	CreatedAt    time.Time  `json:"createdAt"`
	LastLoginAt  *time.Time `json:"lastLoginAt,omitempty"`
}
