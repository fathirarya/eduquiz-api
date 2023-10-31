package domain

import "time"

type Quiz struct {
	ID             uint
	CreatedAt      time.Time
	UpdatedAt      time.Time
	TeacherID      uint
	Teacher        Teacher
	QuizCategoryID uint
	QuizCategory   QuizCategory
	Title          string
	Description    string
	Category       string
	FullName       string
}
