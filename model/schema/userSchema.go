package schema

import (
	"time"

	"gorm.io/gorm"
)

type Users struct {
	ID        uint           `gorm:"primaryKey"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime:milli"`
	DeleteAt  gorm.DeletedAt `gorm:"index"`
	Username  string
	Password  string
	Email     string
	Fullname  string
	Roles     string `gorm:"type:ENUM('siswa', 'guru');not null;default:'siswa'"`
}
