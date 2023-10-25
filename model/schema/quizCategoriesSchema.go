package schema

import (
	"gorm.io/gorm"
)

type QuizCategory struct {
	ID       uint           `gorm:"primarykey"`
	DeleteAt gorm.DeletedAt `gorm:"index"`
	Category string         `gorm:"type:varchar(255);not null"`
}
