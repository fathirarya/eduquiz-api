package schema

import (
	"gorm.io/gorm"
)

type QuizCategory struct {
	ID       uint           `json:"id" gorm:"primarykey" `
	DeleteAt gorm.DeletedAt `gorm:"index"`
	Category string         `json:"category" gorm:"type:varchar(255);not null"`
}
