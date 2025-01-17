package models

import (
	"time"

	"github.com/google/uuid"
)

type EmailCode struct {
	ID               uuid.UUID `db:"id"`
	Email            string    `db:"email"`
	Code             uint16    `db:"code"`
	ExpiresAt        time.Time `db:"expires_at"`
	NumberOfAttempts uint8     `db:"number_of_attempts"`
}
