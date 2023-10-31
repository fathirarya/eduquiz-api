package domain

type AttemptAnswer struct {
	ID         uint
	StudentID  uint
	Student    Student
	QuizID     uint
	Quiz       Quiz
	QuestionID uint
	Question   Question
	Answer     string
	IsCorrect  bool
}
