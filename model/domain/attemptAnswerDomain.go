package domain

type AttemptAnswer struct {
	ID         uint
	QuestionID uint
	Question   Question
	Answer     string
	IsCorrect  bool
	// Quest      string
}
