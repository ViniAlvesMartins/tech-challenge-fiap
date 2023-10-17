package entity

import (
	"time"
)

type User struct {
	ID        uint
	Firstname string
	Lastname  string
	Age       int
	CreatedAt time.Time
	UpdatedAt time.Time
}
