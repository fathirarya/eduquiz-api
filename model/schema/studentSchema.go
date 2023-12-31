package schema

import (
	"time"

	"gorm.io/gorm"
)

type Student struct {
	ID            uint           `gorm:"primaryKey"`
	CreatedAt     time.Time      `gorm:"autoCreateTime"`
	UpdatedAt     time.Time      `gorm:"autoUpdateTime:milli"`
	DeleteAt      gorm.DeletedAt `gorm:"index"`
	Username      string
	Password      string
	Email         string
	Fullname      string
	AttemptAnswer []AttemptAnswer `gorm:"foreignKey:StudentID"`
}
