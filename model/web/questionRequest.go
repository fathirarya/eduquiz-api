package web

type QuestionCreateRequest struct {
	QuizID   uint   `json:"quiz_id" validate:"required"`
	Question string `json:"question" validate:"required,min=1"`
	Option1  string `json:"option_1" validate:"required,min=1"`
	Option2  string `json:"option_2" validate:"required,min=1"`
	Option3  string `json:"option_3" validate:"required,min=1"`
	Option4  string `json:"option_4" validate:"required,min=1"`
}

type QuestionUpdateRequest struct {
	QuizID   uint   `json:"quiz_id" validate:"required"`
	Question string `json:"question" validate:"required,min=1"`
	Option1  string `json:"option_1" validate:"required,min=1"`
	Option2  string `json:"option_2" validate:"required,min=1"`
	Option3  string `json:"option_3" validate:"required,min=1"`
	Option4  string `json:"option_4" validate:"required,min=1"`
}
