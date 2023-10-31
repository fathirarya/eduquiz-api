package schema

import "gorm.io/gorm"

type QuizResult struct {
	gorm.Model
	StudentID uint    `gorm:"index"`
	Student   Student `gorm:"foreignkey:StudentID"`
	QuizID    uint    `gorm:"index"`
	Quiz      Quiz    `gorm:"foreignkey:QuizID"`
	Score     int     `gorm:"type:int;not null"`
}
