package schema

import (
	"gorm.io/gorm"
)

type Quiz struct {
	gorm.Model
	QuizCategoryID uint         `gorm:"index"`
	QuizCategory   QuizCategory `gorm:"foreignkey:QuizCategoryID"`
	TeacherID      uint         `gorm:"index"`
	Teacher        Teacher      `gorm:"foreignkey:TeacherID"`
	Title          string       `gorm:"type:varchar(255);not null"`
	Description    string       `gorm:"type:varchar(255);not null"`
}
