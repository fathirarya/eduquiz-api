package domain

import "time"

type Teacher struct {
	ID        uint
	CreatedAt time.Time
	UpdatedAt time.Time
	Username  string
	Password  string
	Email     string
	Fullname  string
}
