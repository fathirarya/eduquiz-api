package schema

import (
	"time"

	"gorm.io/gorm"
)

type Users struct {
	ID        uint           `gorm:"primary_key"`
	CreatedAt time.Time      `gorm:"created_at"`
	UpdatedAt time.Time      `gorm:"updated_at:milli"`
	DeleteAt  gorm.DeletedAt `gorm:"index"`
	Username  string
	Password  string
	Email     string
	Fullname  string
	Roles     string `gorm:"type:ENUM('siswa', 'guru');not null;default:'siswa'"`
}
