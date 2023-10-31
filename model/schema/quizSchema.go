package schema

import (
	"gorm.io/gorm"
)

type Quiz struct {
	gorm.Model
	QuizCategoryID uint         `gorm:"index"`
	QuizCategory   QuizCategory `gorm:"foreignkey:QuizCategoryID"`
	Title          string       `gorm:"type:varchar(255);not null"`
	Description    string       `gorm:"type:varchar(255);not null"`
	QuizResults    []QuizResult `gorm:"foreignKey:QuizID"`
}
