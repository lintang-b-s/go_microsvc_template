package jwt

import (
	"errors"
	"time"
)

var (
	ErrInvalidToken = errors.New("token is invalid")
	ErrExpiredToken = errors.New("token has expired")
)

// Payload berisi payload data dari token
type Payload struct {
	ID        string    `json:"id"`
	Username  string    `json:"username"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"exp"`
}

// NewPayload membuat payload token baru
func NewPayload(userID, username string, duration time.Duration) (*Payload, error) {
	// tokenID, err := uuid.NewRandom()
	// if err != nil {
	// 	return nil, err
	// }
	// userIDUUID, _ := gofrsUUID.FromString(userID)

	payload := &Payload{
		ID:        userID,
		Username:  username,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}
	return payload, nil
}

// Valid cek jika token payload valid atau tidak
func (payload *Payload) Valid() error {
	if time.Now().After(payload.ExpiredAt) {
		return ErrExpiredToken
	}
	return nil
}
