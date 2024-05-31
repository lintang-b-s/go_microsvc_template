package domain

import (
	"time"
)

type Session struct {
	ID           string    `json:"id"`
	Username     string    `json:"username"`
	RefreshToken string    `json:"refresh_token"`
	ExpiresAt    time.Time `json:"expires_at"`
}


