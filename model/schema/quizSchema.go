package schema

import (
	"gorm.io/gorm"
)

type Quiz struct {
	gorm.Model
	Title          string `gorm:"type:varchar(255);not null"`
	Description    string `gorm:"type:varchar(255);not null"`
	QuizCategoryID uint   `gorm:"index"`
	QuizCategory   uint   `gorm:"foreignkey:QuizCategoryID"`
}
