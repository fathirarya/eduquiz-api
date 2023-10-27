package schema

import "gorm.io/gorm"

type Question struct {
	gorm.Model
	QuizID   uint   `gorm:"index"`
	Quiz     Quiz   `gorm:"foreignkey:QuizID"`
	Question string `gorm:"type:text;not null"`
	Option1  string `gorm:"type:text;not null"`
	Option2  string `gorm:"type:text;not null"`
	Option3  string `gorm:"type:text;not null"`
	Option4  string `gorm:"type:text;not null"`
}
