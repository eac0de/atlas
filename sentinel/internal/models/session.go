package models

import (
	"net"
	"time"

	"github.com/google/uuid"
)

type Session struct {
	ID         uuid.UUID `db:"id" json:"id"`
	Token      string    `db:"token" json:"-"`
	UserID     uuid.UUID `db:"user_id" json:"-"`
	IP         net.IP    `db:"ip" json:"-"`
	Location   string    `db:"location" json:"location"`
	ClientInfo string    `db:"client_info" json:"client_info"`
	LastLogin  time.Time `db:"last_login" json:"last_login"`
}
