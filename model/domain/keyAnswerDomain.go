package domain

type KeyAnswer struct {
	ID         uint
	QuestionID uint
	Question   Question
	Answer     string
	Questions  string
}
