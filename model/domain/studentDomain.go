package domain

import "time"

type Student struct {
	ID        uint
	CreatedAt time.Time
	UpdatedAt time.Time
	DeleteAt  time.Time
	Username  string
	Password  string
	Email     string
	Fullname  string
}
