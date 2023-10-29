package domain

type Quiz struct {
	ID             uint
	QuizCategoryID uint
	QuizCategory   QuizCategory
	Title          string
	Description    string
	Category       string
}
