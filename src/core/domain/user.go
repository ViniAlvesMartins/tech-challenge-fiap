package domain

import (
	"time"
)

type User struct {
	ID        uint32
	Nickname  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
