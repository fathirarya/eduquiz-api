package schema

type AttemptAnswer struct {
	ID         uint     `json:"id" gorm:"primaryKey"`
	StudentID  uint     `json:"student_id"`
	Student    Student  `json:"student" gorm:"foreignKey:StudentID"`
	QuizID     uint     `json:"quiz_id"`
	Quiz       Quiz     `json:"quiz" gorm:"foreignKey:QuizID"`
	QuestionID uint     `json:"question_id"`
	Question   Question `json:"question" gorm:"foreignKey:QuestionID"`
	Answer     string   `json:"answer"`
	IsCorrect  bool     `json:"is_correct"`
}
